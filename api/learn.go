package api

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"LipLanguage/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func UploadVideo(ctx *gin.Context) {
	VideoIDRaw := ctx.PostForm("video_id")
	VideoDataRaw, err := ctx.FormFile("video")
	if err != nil {
		logrus.Errorf("[api.UpdateVideo] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "视频错误",
		})
		return
	}

	VideoID, err := strconv.Atoi(VideoIDRaw)
	if err != nil {
		logrus.Errorf("[api.UpdateVideo] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数非法",
		})
		return
	}

	token := ctx.GetHeader("auth")
	claim, _ := dao.ParseToken(token)

	res, err := service.UploadVideo(ctx, claim.Phone, int64(VideoID), VideoDataRaw)
	if err != nil {
		logrus.Errorf("[api.UpdateVideo] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "上传成功",
			"data": gin.H{
				"result": res,
			},
		})
		return
	}
}

func GetVideoHistory(ctx *gin.Context) {
	LimitRaw := ctx.Query("limit")
	OffsetRaw := ctx.Query("offset")
	VideoIDRaw := ctx.Query("video_id")
	Limit, err1 := strconv.Atoi(LimitRaw)
	Offset, err2 := strconv.Atoi(OffsetRaw)
	VideoID, err3 := strconv.Atoi(VideoIDRaw)
	if err1 != nil || err2 != nil || err3 != nil {
		logrus.Errorf("[api.GetVideoHistory] %v %v %v", err1, err2, err3)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数非法",
		})
		return
	}

	data, err := service.GetVideoHistory(int64(VideoID), Offset, Limit)
	if err != nil {
		logrus.Errorf("[api.GetVideoHistory] %v %v %v", err1, err2, err3)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "请求成功",
		"data": data,
	})
}

func GetTodayRecord(ctx *gin.Context) {
	token := ctx.GetHeader("auth")
	claim, _ := dao.ParseToken(token)
	data, err := service.GetTodayLearnData(claim.Phone)
	if err != nil {
		logrus.Errorf("[api.GetTodayRecord] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "请求成功",
		"data": data,
	})

}

func GetAllStandardVideos(ctx *gin.Context) {
	LimitRaw := ctx.Query("limit")
	OffsetRaw := ctx.Query("offset")
	Limit, err1 := strconv.Atoi(LimitRaw)
	Offset, err2 := strconv.Atoi(OffsetRaw)
	if err1 != nil || err2 != nil {
		logrus.Errorf("[api.GetVideoHistory] %v %v", err1, err2)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数非法",
		})
		return
	}

	data, err := service.GetAllStandardVideos(Limit, Offset)
	if err != nil {
		logrus.Errorf("[api.GetAllStandardVideos] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		return
	} else {
		if data == nil {
			data = make([]model.StandardVideo, 0)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "请求成功",
			"data": data,
		})
	}
}

func GetDayHistory(ctx *gin.Context) {
	LimitRaw := ctx.Query("limit")
	OffsetRaw := ctx.Query("offset")
	Limit, err1 := strconv.Atoi(LimitRaw)
	Offset, err2 := strconv.Atoi(OffsetRaw)
	if err1 != nil || err2 != nil {
		logrus.Errorf("[api.GetVideoHistory] %v %v", err1, err2)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数非法",
		})
		return
	}
	token := ctx.GetHeader("auth")
	claim, _ := dao.ParseToken(token)
	data, err := service.GetDayHistory(Limit, Offset, claim.UserID)
	if err != nil {
		logrus.Errorf("[api.GetAllStandardVideos] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		return
	} else {
		if data == nil {
			data = make([]model.LearnStatistics, 0)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "请求成功",
			"data": data,
		})
	}
}
