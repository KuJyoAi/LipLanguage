package user

import (
	"LipLanguage/dao"
	"LipLanguage/model"
)

func GetNotice(UserID int64, limit int, offset int) (notice []model.Notice, err error) {
	err = dao.DB.
		Where("user_id = ?", UserID).
		Limit(limit).Offset(offset).
		Find(&notice).Error
	return
}
