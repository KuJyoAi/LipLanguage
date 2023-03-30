package midware

import (
	"LipLanguage/dao"
	"LipLanguage/dao/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func FromReqGetToken(ctx *gin.Context) string {
	token := ctx.PostForm("auth")
	return token
}
func FromReqGetClaims(ctx *gin.Context) dao.Claim {
	token := ctx.PostForm("auth")
	c, err := dao.ParseToken(token)
	if err != nil {
		logrus.Infof("[midware] ParseToken %v", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":  "token无效",
			"data": gin.H{},
		})
	}
	return *c
}

func Auth(ctx *gin.Context) {
	logrus.Infof("[midware] Auth %v", ctx.Request.URL.Path)
	logrus.Infof("[midware] Cookies %v", ctx.Request.Cookies())
	token := ctx.PostForm("auth")

	if token == "" {
		logrus.Infof("[midware] Cookie %v", token)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":  "No Token!",
			"data": gin.H{},
		})
		return
	}

	//解析token
	claims, err := dao.ParseToken(token)
	if err != nil {
		logrus.Infof("[midware] ParseToken %v", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":  "token无效",
			"data": gin.H{},
		})
		return
	}

	//检查用户是否存在
	if !user.Exists(claims.Phone) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "token无效",
		})
		return
	}

	//token是否过期
	key := fmt.Sprintf("%v_Token", claims.Phone)
	RedisToken, err := dao.RDB.Get(key).Result()
	if err != nil || RedisToken != token {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "token无效",
		})
		return
	}

	ctx.Next()
}

// CORS 跨域问题
func CORS(context *gin.Context) {
	logrus.Infof("[midware] CORS: Path=%v, Origin=%v", context.Request.URL.Path, context.GetHeader("Origin"))
	method := context.Request.Method
	// 1. [必须]接受指定域的请求，可以使用*不加以限制，但不安全
	//context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
	//fmt.Println(context.GetHeader("Origin"))
	// 2. [必须]设置服务器支持的所有跨域请求的方法
	context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// 3. [可选]服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Cookie")
	// 4. [可选]设置XMLHttpRequest的响应对象能拿到的额外字段
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Cookie")
	// 5. [可选]是否允许后续请求携带认证信息Cookie，该值只能是true，不需要则不设置
	context.Header("Access-Control-Allow-Credentials", "true")
	// 6. 放行所有OPTIONS方法
	if method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
	}
	context.Next()
}
