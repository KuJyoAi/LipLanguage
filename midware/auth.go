package midware

import (
	"jcz-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Auth(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		token = ctx.PostForm("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "No Token!",
			})
		}
		ctx.Abort()
		return
	}

	//解析token
	claims, err := utils.VerifyJWT(token)
	if err != nil {
		logrus.Infof("[midware] VerifyJWT %v", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "token无效",
		})
		ctx.Abort()
		return
	}

	ctx.Set("user_id", claims.UserID)
	ctx.Set("user_phone", claims.Phone)
	ctx.Set("user_nickname", claims.Nickname)

	ctx.Next()
}
