package api

import (
	"LipLanguage/midware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://www.jczlipread.cn",
			"http://localhost:3000",
			"https://jczlipread.cn"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Cookie", "X-Requested-With", "auth"},
		ExposeHeaders:    []string{"Content-Length", "Cookie", "auth"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour})) // 处理跨域问题

	api := r.Group("api")
	{
		user := api.Group("user")
		{
			user.POST("/register", Register)
			user.POST("/login", Login)
			user.POST("/update", midware.Auth, UserInfoUpdate)
			user.POST("/verify", UserVerify)
			user.POST("/resetpassword", ResetPassword)
			user.POST("/updatephone", midware.Auth, UserUpdatePhone)
			user.POST("/updatepassword", midware.Auth, UserUpdatePassword)
			user.POST("/profile", midware.Auth, UserProfile)
			user.POST("/getnotice", midware.Auth, UserGetNotice)
			user.POST("/readnotice", midware.Auth, UserReadNotice)
		}

		learn := api.Group("learn")
		learn.Use(midware.Auth)
		{
			learn.POST("/getStandards", GetStandardVideos)
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

		api.POST("/resource", midware.Auth, GetResource)
	}

	// 测试页面
	r.GET("/", ReturnIndex)
	r.LoadHTMLGlob("html/*")
}
