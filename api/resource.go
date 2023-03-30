package api

import (
	"LipLanguage/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func GetResource(ctx *gin.Context) {
	SrcID := ctx.PostForm("src_id")
	logrus.Infof("GetResource %v", SrcID)
	if SrcID == "" {
		Response(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	data, err := service.GetResource(SrcID)
	if err != nil {
		logrus.Errorf("[api.GetResource]%v", err)
		if err == gorm.ErrRecordNotFound {
			Response(ctx, http.StatusNotFound, "资源不存在", data)
			return
		} else {
			Response(ctx, http.StatusInternalServerError, "服务端错误", data)
			return
		}
	}
	Response(ctx, http.StatusOK, "获取成功", data)
}
