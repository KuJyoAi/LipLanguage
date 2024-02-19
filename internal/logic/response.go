package logic

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, code int, msg string, data interface{}) {
	if data == nil {
		ctx.JSON(code, gin.H{
			"msg": msg,
		})
	}
	ctx.JSON(code, gin.H{
		"msg":  msg,
		"data": data,
	})
}