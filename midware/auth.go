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
		token = ctx.PostForm("auth")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "No Token!",
			})
			return
		}
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
func Access(context *gin.Context) {
	method := context.Request.Method
	// 1. [必须]接受指定域的请求，可以使用*不加以限制，但不安全
	//context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Origin", "*")
	fmt.Println(context.GetHeader("Origin"))
	// 2. [必须]设置服务器支持的所有跨域请求的方法
	context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// 3. [可选]服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, auth")
	// 4. [可选]设置XMLHttpRequest的响应对象能拿到的额外字段
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, auth")
	// 5. [可选]是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	context.Header("Access-Control-Allow-Credentials", "true")
	// 6. 放行所有OPTIONS方法
	if method == "OPTIONS" {
		context.AbortWithStatus(http.StatusOK)
	}
	context.Next()
}
