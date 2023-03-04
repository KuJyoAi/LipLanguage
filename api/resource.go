package api

import (
	"LipLanguage/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetResource(ctx *gin.Context) {
	RawSrcID := ctx.GetHeader("id")
	SrcID, err := strconv.Atoi(RawSrcID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数不合法",
		})
		return
	}

	data, err := service.GetResource(uint(SrcID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		return
	}

	_, err = ctx.Writer.Write(*data)
	if err != nil {
		logrus.Errorf("[api.GetResource]%v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		return
	}
}
