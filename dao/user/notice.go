package user

import (
	"jcz-backend/dao"
	"jcz-backend/model"
)

func GetNotice(UserID int64, limit int, offset int, read string) (notice []model.Notice, err error) {
	if read == "false" {
		err = dao.DB.
			Where("user_id = ? AND status = ?", UserID, 0).
			Limit(limit).Offset(offset).
			Order("created_at DESC").
			Find(&notice).Error
		return
	} else {
		err = dao.DB.
			Where("user_id = ?", UserID).
			Limit(limit).Offset(offset).
			Order("created_at DESC").
			Find(&notice).Error
		return
	}
}

func ReadNotice(UserID int64, noticeID int) (err error) {
	notice := model.Notice{}
	err = dao.DB.
		Model(&model.Notice{}).
		Where("user_id = ? AND id = ?", UserID, noticeID).
		First(&notice).Error
	if err != nil {
		return
	}
	notice.Status = 1
	err = dao.DB.Save(&notice).Error
	return
}
