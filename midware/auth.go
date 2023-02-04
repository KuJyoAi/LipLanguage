package midware

import (
	"LipLanguage/dao"
	"LipLanguage/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Auth(ctx *gin.Context) {
	token := ctx.GetHeader("auth")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "No Token!",
		})
		return
	}

	claims, err := util.ParseToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "token无效",
		})
		return
	}
	//检查用户是否存在, token是否过期
	if claims.ExpiresAt < time.Now().Unix() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "token过期",
		})
		return
	}
	if !dao.UserExists(int64(claims.ID)) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "token无效",
		})
		return
	}
}
