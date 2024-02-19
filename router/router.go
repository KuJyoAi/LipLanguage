package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"jcz-backend/internal/logic"
	"jcz-backend/midware"
)

func NewServer() *gin.Engine {
	r := gin.Default()
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

		learn := api.Group("learn", midware.Auth)
		{
			learn.POST("/time", logic.UpdateLearnTime)
			learn.GET("/time", logic.GetLearnTime)

			// TODO
			//learn.POST("/getStandards", logic.GetStandardVideos)
			//learn.POST("/standardHistory", logic.GetStandardVideoLearnHistory)
			//learn.POST("/train", logic.UploadTrainVideo)

		}

		api.POST("/oss/:id", midware.Auth, logic.GetOss)
	}

	return r
}