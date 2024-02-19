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
			Response(c, http.StatusInternalServerError, "内部错误", nil)
			return
		}
		Response(c, http.StatusOK, "请求成功", gin.H{
			"add_time": 0,
		})
	} else if err != nil {
		logrus.Errorf("[api.UpdateLearnTime] %v", err)
		Response(c, http.StatusInternalServerError, "内部错误", nil)
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
			Response(c, http.StatusInternalServerError, "数据库错误", nil)
			return err
		}
		learnTime.LearnTime += time.Now().Unix() - lastTime
		if err = tx.Save(&learnTime).Error; err != nil {
			logrus.Errorf("[api.UpdateLearnTime] %v", err)
			Response(c, http.StatusInternalServerError, "数据库错误", nil)
			return err
		}

		if err = engine.GetRedisCli().Set(fmt.Sprintf("last_learn_time_%d", userID), nowTime, 2*time.Minute).Err(); err != nil {
			logrus.Errorf("[api.UpdateLearnTime] %v", err)
			Response(c, http.StatusInternalServerError, "内部错误", nil)
			return err
		}

		Response(c, http.StatusOK, "请求成功", gin.H{
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
		Response(c, http.StatusBadRequest, "参数错误", nil)
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
		Response(c, http.StatusInternalServerError, "数据库错误", nil)
		return
	}

	Response(c, http.StatusOK, "请求成功", learnTime)
}
