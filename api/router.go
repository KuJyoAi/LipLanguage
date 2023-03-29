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
			user.GET("/profile", midware.Auth, UserProfile)
		}

		learn := api.Group("learn")
		learn.Use(midware.Auth)
		{
			learn.GET("/getStandards", GetStandardVideos)
			learn.POST("/standardHistory", GetStandardVideoLearnHistory)
			learn.POST("/train", UploadTrainVideo)

			learnData := learn.Group("statistics")
			{
				learnData.POST("/today", GetTodayStatistic)
				learnData.POST("/month", GetMonthRecord)
			}
		}
		manager := api.Group("manager")
		{
			manager.POST("/uploadStandard", UploadStandardVideo)
		}

		api.GET("/resource", midware.Auth, GetResource)
	}

	// 测试页面
	r.GET("/", ReturnIndex)
	r.LoadHTMLGlob("html/*")
}
