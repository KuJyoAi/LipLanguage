package api

import (
	"LipLanguage/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func GetResource(ctx *gin.Context) {
	SrcID := ctx.PostForm("src_id")
	if SrcID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数不合法",
		})
		return
	}

	data, err := service.GetResource(SrcID)
	if err != nil {
		logrus.Errorf("[api.GetResource]%v", err)
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "没有此资源",
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "内部错误",
				"data": data,
			})
			return
		}
	}
}
