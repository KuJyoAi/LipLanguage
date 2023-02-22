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
			user.POST("/verify", UserVerify)
			user.POST("/resetPassword", midware.Auth, ResetPassword)
			user.POST("/updatephone", midware.Auth, UserUpdatePhone)
			user.POST("/updatepassword", midware.Auth, UserUpdatePassword)
			user.GET("/profile", midware.Auth, UserProfile)
		}
		learn := api.Group("learn")
		{
			learn.POST("/train", midware.Auth, UpdateVideo)
			learn.GET("/history", midware.Auth, GetVideoHistory)
			learn.GET("/today", midware.Auth, GetTodayRecord)
			learn.GET("/month", midware.Auth, GetRecordsByMonth)
		}
	}

}
