package midware

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// RouterCount 统计路由调用
func RouterCount(ctx *gin.Context) {
	path := ctx.Request.URL.Path

	token := ctx.GetHeader("auth")
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
}
