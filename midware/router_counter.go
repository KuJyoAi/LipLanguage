package midware

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// RouterCount 统计路由调用
func RouterCount(ctx *gin.Context) {
	path := ctx.Request.URL.Path

	token := ctx.GetHeader("auth")
	if token == "" {
		token = ctx.PostForm("auth")
		if token == "" {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
	}

	claim, err := dao.ParseToken(token)
	if err != nil {
		logrus.Infof("[mid.RouterCount] %v", err)
	}
	router := model.RouterCounter{
		Model:  gorm.Model{},
		UserID: claim.UserID,
		Path:   path,
	}
	dao.AddRouterCounter(router)
	logrus.Infof("RouterCount: User:%v, Path:%v", claim.UserID, path)
	// 添加计时器
	start := time.Now()
	ctx.Next()
	// 计算耗时
	cost := time.Since(start)
	logrus.Infof("RouterCount: User:%v, Path:%v, Cost:%v", claim.UserID, path, cost)
}
