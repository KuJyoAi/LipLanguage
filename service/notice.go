package service

import (
	"LipLanguage/dao/user"
	"LipLanguage/model"
	"github.com/sirupsen/logrus"
)

func GetNotice(UserID int64, limit int, offset int) (notice []model.Notice, err error) {
	notice, err = user.GetNotice(UserID, limit, offset)
	if err != nil {
		logrus.Errorf("[service.GetNotice] GetNotice %v", err)
		return
	}
	return
}
