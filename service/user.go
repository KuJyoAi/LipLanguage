package service

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"LipLanguage/util"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func Register(Phone int64, Password string) (string, error) {
	// todo password 哈希
	user := model.User{
		Model:         gorm.Model{},
		AvatarUrl:     "",
		Phone:         Phone,
		Email:         "",
		Name:          "",
		Password:      util.Hash256(Password),
		HearingDevice: false,
		Gender:        0,
		BirthDay:      time.Date(2050, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	err := dao.Register(&user)
	if err != nil {
		logrus.Errorf("[service] Register %v", err)
		return "", err
	}

	token, err := util.GenerateToken(user.ID, user.Nickname)
	if err != nil {
		logrus.Errorf("[service] Register %v", err)
		return "", err
	}

	return token, err
}

func Login(Phone int64, Password string) (string, error) {
	user, err := dao.GetUserByPhone(Phone)
	if err != nil {
		logrus.Errorf("[service.Login] %v", err)
		return "", err
	}

	if util.Hash256(Password) != user.Password {
		logrus.Errorf("[service.Login] %v", err)
		return "", errors.New("PasswordError")
	}

	token, err := util.GenerateToken(user.ID, user.Nickname)
	if err != nil {
		logrus.Errorf("[service.Login] %v", err)
		return "", err
	}

	return token, err
}

func UserInfoUpdate(token string, info *model.UpdateInfoParam) error {
	claim, err := util.ParseToken(token)
	if err != nil {
		logrus.Errorf("[service.UserInfoUpdate]ParseToken %v", err)
		return err
	}

	UserID := int64(claim.ID)
	err = dao.UserInfoUpdate(UserID, info)
	if err != nil {
		logrus.Errorf("[service.UserInfoUpdate] %v", err)
	}
	return err
}
