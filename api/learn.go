package api

import (
	"LipLanguage/dao"
	"LipLanguage/service/learn"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"time"
)

func UploadTrainVideo(ctx *gin.Context) {
	VideoIDRaw := ctx.PostForm("video_id")
	video, header, err := ctx.Request.FormFile("video")
	if err != nil {
		logrus.Errorf("[api.UpdateVideo] %v", err)
		Response(ctx, http.StatusBadRequest, "视频错误", nil)
		return
	}
	logrus.Infof("[api.UploadVideo] From frontend:\nname:%v size:%v Header:%v\n time=%v",
		header.Filename, header.Size, header.Header, time.Now())
	VideoID, err := strconv.Atoi(VideoIDRaw)

	videoData, err := io.ReadAll(video)
	if err != nil {
		logrus.Errorf("[api.UpdateVideo] %v", err)
		Response(ctx, http.StatusBadRequest, "视频错误", nil)
		return
	}

	token, _ := ctx.Cookie("auth")
	claim, _ := dao.ParseToken(token)

	res, err := learn.UploadTrainVideo(claim.Phone, int64(VideoID), videoData)

	logrus.Infof("[api.UploadVideo] Send to frontend:\n err:%v ok:%v\n time=%v",
		err, time.Now())
	if err != nil {
		Response(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	} else {
		Response(ctx, http.StatusOK, "上传成功", res)
		return
	}
}

func GetTodayStatistic(ctx *gin.Context) {
	token, _ := ctx.Cookie("auth")
	claim, _ := dao.ParseToken(token)

	data, err := learn.GetTodayStatistic(claim.UserID)

	if err != nil {
		logrus.Errorf("[api.GetTodayRecord] %v", err)
		Response(ctx, http.StatusInternalServerError, "内部错误", nil)
		return
	} else {
		Response(ctx, http.StatusOK, "请求成功", data)
	}
}

func GetMonthRecord(ctx *gin.Context) {
	YearRaw := ctx.PostForm("year")
	MonthRaw := ctx.PostForm("month")
	Year, err1 := strconv.Atoi(YearRaw)
	Month, err2 := strconv.Atoi(MonthRaw)
	if err1 != nil || err2 != nil {
		logrus.Errorf("[api.GetMonthRecord] %v \n%v", err1, err2)
		Response(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}
	token, _ := ctx.Cookie("auth")
	claim, _ := dao.ParseToken(token)

	data, err := learn.GetMonthStatistic(claim.UserID, Year, Month)
	
	if err != nil {
		logrus.Errorf("[api.GetMonthRecord] %v", err)
		Response(ctx, http.StatusInternalServerError, "内部错误", nil)
		return
	} else {
		Response(ctx, http.StatusOK, "请求成功", data)
	}
}
func GetStandardVideoLearnHistory(ctx *gin.Context) {
	LimitRaw := ctx.PostForm("limit")
	OffsetRaw := ctx.PostForm("offset")
	VideoIDRaw := ctx.PostForm("video_id")
	Order := ctx.PostForm("order")
	Limit, err1 := strconv.Atoi(LimitRaw)
	Offset, err2 := strconv.Atoi(OffsetRaw)
	VideoID, err3 := strconv.Atoi(VideoIDRaw)
	if err1 != nil || err2 != nil || err3 != nil {
		logrus.Errorf("[api.GetVideoHistory] %v \n%v \n%v", err1, err2, err3)
		Response(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	token, _ := ctx.Cookie("auth")
	claim, _ := dao.ParseToken(token)

	data, err := learn.GetStandardVideoLearnRecord(claim.UserID, int64(VideoID), Limit, Offset, Order)

	if err != nil {
		logrus.Errorf("[api.GetStandardVideoLearnHistory] %v", err)
		Response(ctx, http.StatusInternalServerError, "内部错误", nil)
		return
	} else {
		Response(ctx, http.StatusOK, "请求成功", data)
	}
}

func GetStandardVideos(ctx *gin.Context) {
	LimitRaw := ctx.PostForm("limit")
	OffsetRaw := ctx.PostForm("offset")
	Order := ctx.PostForm("order")
	Limit, err1 := strconv.Atoi(LimitRaw)
	Offset, err2 := strconv.Atoi(OffsetRaw)
	if err1 != nil || err2 != nil {
		logrus.Errorf("[api.GetVideoHistory] %v %v", err1, err2)
		Response(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	data, err := learn.GetAllStandardVideos(Limit, Offset, Order)

	if err != nil {
		logrus.Errorf("[api.GetAllStandardVideos] %v", err)
		Response(ctx, http.StatusInternalServerError, "内部错误", nil)
		return
	} else {
		Response(ctx, http.StatusOK, "请求成功", data)
	}
}
