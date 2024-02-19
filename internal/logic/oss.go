package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jcz-backend/config"
	"os"
)

// GetOss 获取本地存储的资源
func GetOss(ctx *gin.Context) {
	SrcID := ctx.Param("id")

	data, err := os.ReadFile(config.GetConfig().StoragePath + "/" + SrcID)
	if err != nil {
		logrus.Errorf("GetOss %v", err)
		Response(ctx, 404, "资源不存在", nil)
		return
	}

	ctx.Data(200, "application/octet-stream", data)
}
