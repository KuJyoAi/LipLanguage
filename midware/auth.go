package midware

import (
	"LipLanguage/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(ctx *gin.Context) {
	token := ctx.GetHeader("auth")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "No Token!",
		})
		return
	}

	claims, err := dao.ParseToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "token无效",
		})
		return
	}
	//检查用户是否存在
	if !dao.UserExists(claims.Phone) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "token无效",
		})
		return
	}
	//token是否过期
	key := fmt.Sprintf("%v_Token", claims.Phone)
	RedisToken, err := dao.RDB.Get(key).Result()
	if err != nil || RedisToken != token {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "token无效",
		})
		return
	}
}
