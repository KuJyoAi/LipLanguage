package api

import (
	"LipLanguage/midware"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	api := r.Group("api")
	{
		user := api.Group("user")
		{
			user.POST("/register", Register)
			user.POST("/login", Login)
			user.POST("/update", midware.Auth, UserInfoUpdate)
			user.POST("/resetPassword", ResetPassword)
		}
	}

}
