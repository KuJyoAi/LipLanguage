package api

import (
	"LipLanguage/midware"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(midware.CORS) // 处理跨域问题
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
			user.POST("/profile", midware.Auth, UserProfile)
		}
		learn := api.Group("learn")
		{
			learn.POST("/train", midware.Auth, midware.RouterCount, UploadVideo)
			learn.POST("/upload", UploadStandardVideo)
			//learn.GET("/get_svideo", midware.Auth, midware.RouterCount)
			learn.POST("/history", midware.Auth, midware.RouterCount, GetVideoHistory)
			learn.POST("/today", midware.Auth, midware.RouterCount, GetTodayRecord)
			learn.POST("/standard", midware.Auth, midware.RouterCount, GetAllStandardVideos)
			learn.POST("/dayhistory", midware.Auth, midware.RouterCount, GetDayHistory)
		}
		api.POST("/resource", midware.Auth, midware.RouterCount, GetResource)
	}

	// 测试页面
	r.GET("/", ReturnIndex)
	r.LoadHTMLGlob("html/*")
}
