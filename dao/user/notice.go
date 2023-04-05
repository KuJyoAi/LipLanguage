package user

import (
	"LipLanguage/dao"
	"LipLanguage/model"
)

func GetNotice(UserID int64, limit int, offset int, read string) (notice []model.Notice, err error) {
	if read == "true" {
		err = dao.DB.
			Where("user_id = ? AND status = ?", UserID, 0).
			Limit(limit).Offset(offset).
			Order("create_time DESC").
			Find(&notice).Error
		return
	} else {
		err = dao.DB.
			Where("user_id = ?", UserID).
			Limit(limit).Offset(offset).
			Order("create_time DESC").
			Find(&notice).Error
		return
	}
}

func ReadNotice(UserID int64, noticeID int) (err error) {
	err = dao.DB.
		Where("user_id = ? AND id = ?", UserID, noticeID).
		Updates(map[string]interface{}{"status": 0}).Error
	return
}
