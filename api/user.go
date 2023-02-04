package api

import (
	"LipLanguage/model"
	"LipLanguage/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Register(ctx *gin.Context) {
	param := model.RegisterParam{}
	err := ctx.ShouldBindJSON(&param)

	if err != nil {
		logrus.Errorf("[api] Register Bind Json %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	//todo 生成token
	token, err := service.Register(param.Phone, param.Password)

	if err != nil {
		logrus.Errorf("[api.Register] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务端错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func Login(ctx *gin.Context) {
	param := model.LoginParam{}
	err := ctx.ShouldBindJSON(&param)

	if err != nil || (param.Nickname == "" && param.Phone == 0) || param.Password == "" {
		logrus.Errorf("[api.Login] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "登录失败, 参数错误",
		})
		return
	}

	if param.Phone != 0 {
		//手机号登录
		token, err := service.Login(param.Phone, param.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "用户不存在或密码错误",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
			"data": gin.H{
				"token": token,
				"time":  time.Now(),
			},
		})
	} else if param.Nickname != "" {
		//昵称登录
	}
}

func ResetPassword(ctx *gin.Context) {

}

func UserInfoUpdate(ctx *gin.Context) {
	info := &model.UpdateInfoParam{}
	err := ctx.ShouldBindJSON(info)
	if err != nil {
		logrus.Errorf("[api.UserInfoUpdate] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	token := ctx.GetHeader("auth")
	err = service.UserInfoUpdate(token, info)
	if err != nil {
		logrus.Errorf("[api.UserInfoUpdate] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
