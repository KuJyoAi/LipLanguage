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
			user.POST("/resetpassword", midware.Auth, ResetPassword)
			user.POST("/updatephone", midware.Auth, UserUpdatePhone)
			user.POST("/updatepassword", midware.Auth, UserUpdatePassword)
			user.GET("/profile", midware.Auth, UserProfile)
		}
		learn := api.Group("learn")
		{
			learn.POST("/train", midware.Auth, midware.RouterCount, UploadVideo)
			//learn.GET("/stardres", midware.Auth, midware.RouterCount, GetStandardResource)
			learn.GET("/get_svideo", midware.Auth, midware.RouterCount)
			learn.GET("/history", midware.Auth, midware.RouterCount, GetVideoHistory)
			learn.GET("/today", midware.Auth, midware.RouterCount, GetTodayRecord)
			learn.GET("/standard", midware.Auth, midware.RouterCount, GetAllStandardVideos)
			learn.GET("/dayhistory", midware.Auth, midware.RouterCount, GetDayHistory)
		}
		api.GET("/resource", midware.Auth, GetResource)
		//api.POST("/uploadvideo")
	}

}
