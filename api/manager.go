package api

import (
	"LipLanguage/common"
	"LipLanguage/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func UploadStandardVideo(ctx *gin.Context) {
	auth := ctx.GetHeader("auth")
	if auth != common.ManagerAuth {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"msg": "操作失败",
		})
		return
	}
	answer := ctx.Query("answer")
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		return
	}
	err = service.UploadStandardVideo(data, answer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "上传成功",
		})
		return
	}
}
