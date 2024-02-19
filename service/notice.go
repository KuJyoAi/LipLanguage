package service

import (
	"github.com/sirupsen/logrus"
	"jcz-backend/dao/user"
	"jcz-backend/model"
)

func GetNotice(UserID int64, limit int, offset int, read string) (notice []model.Notice, err error) {
	notice, err = user.GetNotice(UserID, limit, offset, read)
	if err != nil {
		logrus.Errorf("[service.GetNotice] GetNotice %v", err)
		return
	}
	return
}

func ReadNotice(UserID int64, noticeID int) (err error) {
	err = user.ReadNotice(UserID, noticeID)
	if err != nil {
		logrus.Errorf("[service.ReadNotice] ReadNotice %v", err)
		return
	}
	return
}
