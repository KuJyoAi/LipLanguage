package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"jcz-backend/internal/logic"
	"jcz-backend/midware"
)

func NewServer() *gin.Engine {
	r := gin.Default()

	//r.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{
	//		"https://www.jczlipread.cn",
	//		"http://localhost:3000",
	//		"https://jczlipread.cn"},
	//	AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Cookie", "X-Requested-With", "auth"},
	//	ExposeHeaders:    []string{"Content-Length", "Cookie", "auth"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return true
	//	},
	//	MaxAge: 12 * time.Hour})) // 处理跨域问题

	r.Use(cors.Default())

	api := r.Group("/api")
	{
		noAuthUser := api.Group("/user")
		{
			noAuthUser.POST("/register", logic.UserRegister)
			noAuthUser.POST("/login", logic.UserLogin)
			noAuthUser.POST("/verify", logic.UserVerify)
		}

		authUser := api.Use(midware.Auth)
		{
			authUser.POST("/update_info", logic.UserInfoUpdate)
			authUser.POST("/update_password", logic.UserUpdatePasswordByToken)
			authUser.GET("/profile", logic.UserProfile)
			authUser.GET("/get_notice", logic.UserGetNotice)
			authUser.POST("/read_notice", logic.UserReadNotice)
		}

		learn := api.Group("learn")
		learn.Use(midware.Auth)
		{
			learn.POST("/getStandards", logic.GetStandardVideos)
			learn.POST("/standardHistory", logic.GetStandardVideoLearnHistory)
			learn.POST("/train", logic.UploadTrainVideo)

			learnData := learn.Group("statistics")
			{
				learnData.POST("/today", logic.GetTodayStatistic)
				learnData.POST("/month", logic.GetMonthRecord)
			}
		}
		manager := api.Group("manager")
		{
			manager.POST("/uploadStandard", logic.UploadStandardVideo)
		}

		api.POST("/resource", midware.Auth, logic.GetResource)
	}

	return r
}
