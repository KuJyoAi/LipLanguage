package api

import (
	"LipLanguage/service"
	"LipLanguage/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

func UpdateVideo(ctx *gin.Context) {
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logrus.Errorf("[api.UpdateVideo] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "视频错误",
		})
	}

	VideoID, err := strconv.Atoi(ctx.GetHeader("video_id"))
	if err != nil {
		logrus.Errorf("[api.UpdateVideo] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数非法",
		})
	}

	token := ctx.GetHeader("auth")
	claim, err := util.ParseToken(token)
	res, err := service.UpdateVideo(claim.Phone, int64(VideoID), &data)
	if err != nil {
		logrus.Errorf("[api.UpdateVideo] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "上传成功",
			"data": gin.H{
				"result": res,
			},
		})
	}
}

func GetVideoHistory(ctx *gin.Context) {

}

func GetTodayRecord(ctx *gin.Context) {

}

func GetRecordsByMonth(ctx *gin.Context) {

}
