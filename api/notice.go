package api

import (
	"LipLanguage/midware"
	"LipLanguage/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func UserGetNotice(ctx *gin.Context) {
	limitStr := ctx.PostForm("limit")
	offsetStr := ctx.PostForm("offset")
	limit, err1 := strconv.Atoi(limitStr)
	offset, err2 := strconv.Atoi(offsetStr)
	if err1 != nil || err2 != nil {
		logrus.Errorf("[api.UserGetNotice] %v %v", err1, err2)
		Response(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}

	claim := midware.FromReqGetClaims(ctx)
	notice, err := service.GetNotice(claim.UserID, limit, offset)
	if err != nil {
		logrus.Errorf("[api.UserGetNotice] %v", err)
		Response(ctx, http.StatusInternalServerError, "服务端错误", nil)
		return
	} else {
		Response(ctx, http.StatusOK, "获取成功", notice)
	}
}
