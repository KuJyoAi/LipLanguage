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

// Access 跨域问题
func Access(ctx *gin.Context) {
	method := ctx.Request.Method
	ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
	// 设置服务器支持的跨域请求方法
	ctx.Header("Access-Control-Allow-Methods", "POST, GET")
	// 设置服务器支持的头信息字段
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}
