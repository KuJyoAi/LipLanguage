package logic

import (
	"github.com/gin-gonic/gin"
	"jcz-backend/internal/engine"
	"jcz-backend/model"
	"net/http"
)

func UserGetNotice(ctx *gin.Context) {
	type Params struct {
		PageIdx  int  `json:"page_idx" form:"page_idx"` // 从1开始
		PageSize int  `json:"page_size" form:"page_size"`
		Read     bool `json:"read" form:"read"` // 是否已读
	}

	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	if params.PageIdx < 1 || params.PageSize < 0 || params.PageSize > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	limit := params.PageSize
	offset := (params.PageIdx - 1) * params.PageSize

	userID := ctx.GetUint("user_id")
	var notices []model.Notice
	if err := engine.GetSqlCli().Model(model.Notice{}).
		Where("user_id = ?", userID).
		Where("read = ?", params.Read).
		Order("created_at desc").
		Offset(offset).
		Limit(limit).
		Find(&notices).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据库查询错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": notices,
	})
}

func UserReadNotice(ctx *gin.Context) {
	type Params struct {
		IDs []uint `json:"ids" form:"ids"`
	}

	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	userID := ctx.GetUint("user_id")
	if err := engine.GetSqlCli().Model(model.Notice{}).
		Where("user_id = ?", userID).
		Where("id in ?", params.IDs).
		Update("read", true).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据库更新错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
