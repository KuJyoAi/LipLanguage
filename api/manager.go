package api

import (
	"LipLanguage/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// UploadStandardVideo 上传标准视频
func UploadStandardVideo(ctx *gin.Context) {
	answer := ctx.PostForm("answer")
	video, err1 := ctx.FormFile("video")
	lip, err2 := ctx.FormFile("lip_video")
	if err1 != nil || err2 != nil {
		logrus.Infof("[api.UploadStandardVideo] answer=%v err1=%v err2=%v",
			answer, err1, err2)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "参数错误",
		})
		return
	}

	videoId, lipId, err := service.UploadStandardVideo(ctx, video, lip, answer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "上传失败",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   "上传成功",
			"video": videoId,
			"lip":   lipId,
		})
		return
	}
}
