package logic

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"jcz-backend/internal/engine"
	"jcz-backend/model"
	"jcz-backend/utils"
	"net/http"
	"time"
)

func UpdateLearnTime(c *gin.Context) {
	userID := c.GetUint("user_id")
	lastTime, err := engine.GetRedisCli().Get(fmt.Sprintf("last_learn_time_%d", userID)).Int64()
	nowTime := time.Now().Unix()
	if errors.Is(err, redis.Nil) {
		// 没有学习记录, 直接更新时间
		if err := engine.GetRedisCli().Set(fmt.Sprintf("last_learn_time_%d", userID), nowTime, 2*time.Minute).Err(); err != nil {
			logrus.Errorf("[api.UpdateLearnTime] %v", err)
			utils.Response(c, http.StatusInternalServerError, "内部错误", nil)
			return
		}
		utils.Response(c, http.StatusOK, "请求成功", gin.H{
			"add_time": 0,
		})
	} else if err != nil {
		logrus.Errorf("[api.UpdateLearnTime] %v", err)
		utils.Response(c, http.StatusInternalServerError, "内部错误", nil)
		return
	}

	// 2分钟内有学习记录, 时间累加, 并且更新时间
	_ = engine.GetSqlCli().Transaction(func(tx *gorm.DB) error {
		var learnTime model.UserLearnTime
		todayInt := time.Now().Format("20060102")
		err = tx.Model(&model.UserLearnTime{}).Where("user_id = ? and time_int = ?", userID, todayInt).
			FirstOrCreate(&learnTime).Error
		if err != nil {
			logrus.Errorf("[api.UpdateLearnTime] %v", err)
			utils.Response(c, http.StatusInternalServerError, "数据库错误", nil)
			return err
		}
		learnTime.LearnTime += time.Now().Unix() - lastTime
		if err = tx.Save(&learnTime).Error; err != nil {
			logrus.Errorf("[api.UpdateLearnTime] %v", err)
			utils.Response(c, http.StatusInternalServerError, "数据库错误", nil)
			return err
		}

		if err = engine.GetRedisCli().Set(fmt.Sprintf("last_learn_time_%d", userID), nowTime, 2*time.Minute).Err(); err != nil {
			logrus.Errorf("[api.UpdateLearnTime] %v", err)
			utils.Response(c, http.StatusInternalServerError, "内部错误", nil)
			return err
		}

		utils.Response(c, http.StatusOK, "请求成功", gin.H{
			"add_time": time.Now().Unix() - lastTime,
		})
		return nil
	})
}

func GetLearnTime(c *gin.Context) {
	userID := c.GetUint("user_id")

	type LearnTimeRequest struct {
		PageIdx  int `json:"page_idx" form:"page_idx" binding:"required"`
		PageSize int `json:"page_size" form:"page_size" binding:"required"`
	}

	var req LearnTimeRequest
	if err := c.ShouldBind(&req); err != nil || req.PageIdx < 1 || req.PageSize < 0 {
		logrus.Errorf("[api.GetLearnTime] %v", err)
		utils.Response(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	limit := req.PageSize
	offset := (req.PageIdx - 1) * req.PageSize

	sqlCli := engine.GetSqlCli()
	var learnTime []model.UserLearnTime
	err := sqlCli.Model(&model.UserLearnTime{}).Where("user_id = ?", userID).Order("time_int desc").
		Limit(limit).Offset(offset).Find(&learnTime).Error
	if err != nil {
		logrus.Errorf("[api.GetLearnTime] %v", err)
		utils.Response(c, http.StatusInternalServerError, "数据库错误", nil)
		return
	}

	utils.Response(c, http.StatusOK, "请求成功", learnTime)
}

func GetQuestions(c *gin.Context) {
	type Request struct {
		PageIdx  int `json:"page_idx" form:"page_idx" binding:"required"`
		PageSize int `json:"page_size" form:"page_size" binding:"required"`
	}

	var req Request
	if err := c.ShouldBind(&req); err != nil || req.PageIdx < 1 || req.PageSize < 0 {
		logrus.Errorf("[api.GetQuestions] %v", err)
		utils.Response(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	limit := req.PageSize
	offset := (req.PageIdx - 1) * req.PageSize

	sqlCli := engine.GetSqlCli()
	var questions []model.Question
	err := sqlCli.Model(&model.Question{}).Order("id desc").Limit(limit).Offset(offset).Find(&questions).Error
	if err != nil {
		logrus.Errorf("[api.GetQuestions] %v", err)
		utils.Response(c, http.StatusInternalServerError, "数据库错误", nil)
		return
	}

	utils.Response(c, http.StatusOK, "请求成功", questions)
}

func AnswerQuestion(c *gin.Context) {
	type AnswerQuestionRequest struct {
		QuestionID uint   `json:"question_id" form:"question_id" binding:"required"`
		Result     string `json:"result" form:"result" binding:"required"`
	}
	type AnswerQuestionResponse struct {
		Right      bool   `json:"right"`
		Answer     string `json:"answer"`
		UserAnswer string `json:"user_answer"`
	}

	var req AnswerQuestionRequest
	if err := c.ShouldBind(&req); err != nil {
		logrus.Errorf("[api.AnswerQuestion] %v", err)
		utils.Response(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	userID := c.GetUint("user_id")

	// judge the answer
	sqlCli := engine.GetSqlCli()
	var question model.Question
	err := sqlCli.Model(&model.Question{}).Where("id = ?", req.QuestionID).First(&question).Error
	if err != nil {
		logrus.Errorf("[api.AnswerQuestion] %v", err)
		utils.Response(c, http.StatusInternalServerError, "数据库错误", nil)
		return
	}

	right := question.Answer == req.Result
	learnRecord := model.UserLearnRecord{
		UserID:     userID,
		QuestionID: req.QuestionID,
		Result:     req.Result,
		Right:      right,
	}

	err = sqlCli.Create(&learnRecord).Error
	if err != nil {
		logrus.Errorf("[api.AnswerQuestion] %v", err)
		utils.Response(c, http.StatusInternalServerError, "数据库错误", nil)
		return
	}

	utils.Response(c, http.StatusOK, "请求成功", AnswerQuestionResponse{
		Right:  right,
		Answer: question.Answer,
	})
}

func GetLearnHistory(c *gin.Context) {
	userID := c.GetUint("user_id")

	type LearnHistoryRequest struct {
		PageIdx  int `json:"page_idx" form:"page_idx" binding:"required"`
		PageSize int `json:"page_size" form:"page_size" binding:"required"`
	}

	var req LearnHistoryRequest
	if err := c.ShouldBind(&req); err != nil || req.PageIdx < 1 || req.PageSize < 0 {
		logrus.Errorf("[api.GetLearnHistory] %v", err)
		utils.Response(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	limit := req.PageSize
	offset := (req.PageIdx - 1) * req.PageSize

	sqlCli := engine.GetSqlCli()
	var learnHistory []model.UserLearnRecord
	err := sqlCli.Model(&model.UserLearnRecord{}).Where("user_id = ?", userID).Order("id desc").
		Limit(limit).Offset(offset).Find(&learnHistory).Error
	if err != nil {
		logrus.Errorf("[api.GetLearnHistory] %v", err)
		utils.Response(c, http.StatusInternalServerError, "数据库错误", nil)
		return
	}

	utils.Response(c, http.StatusOK, "请求成功", learnHistory)
}
