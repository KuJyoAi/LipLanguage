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

	token, err := util.GenerateToken(user.Phone, user.Nickname)
	if err != nil {
		logrus.Errorf("[service] Register %v", err)
		return "", err
	}

	return token, err
}

func LoginByPhone(Phone int64, Password string) (string, error) {
	user, err := dao.GetUserByPhone(Phone)
	if err != nil {
		logrus.Errorf("[service.Login] %v", err)
		return "", err
	}

	if util.Hash256(Password) != user.Password {
		logrus.Errorf("[service.Login] %v", err)
		return "", errors.New("PasswordError")
	}

	token, err := util.GenerateToken(user.Phone, user.Nickname)
	if err != nil {
		logrus.Errorf("[service.Login] %v", err)
		return "", err
	}

	return token, err
}

func LoginByNickname(Nickname string, Password string) (string, error) {
	user, err := dao.GetUserByNickname(Nickname)
	if err != nil {
		logrus.Errorf("[service.Login] %v", err)
		return "", err
	}

	if util.Hash256(Password) != user.Password {
		logrus.Errorf("[service.Login] %v", err)
		return "", errors.New("PasswordError")
	}

	token, err := util.GenerateToken(user.Phone, user.Nickname)
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

	err = dao.UserInfoUpdate(claim.Phone, info)
	if err != nil {
		logrus.Errorf("[service.UserInfoUpdate] %v", err)
	}
	return err
}

func UserVerify(Phone int64, Name string, Email string) (string, bool) {
	user, err := dao.GetUserByPhone(Phone)
	if err != nil {
		logrus.Errorf("[service.UserVerify] %v", err)
		return "", false
	}
	if user.Email == Email && user.Name == Name {
		//生成验证码, 5分钟有效
		token, err := util.GenerateTokenExpires(user.Phone, user.Nickname, 5*time.Minute)
		if err != nil {
			logrus.Errorf("[service.UserVerify] %v", err)
			return "", false
		}
		return token, true
	} else {
		return "", false
	}
}

func UserResetPassword(Phone int64, Password string) bool {
	err := dao.UserResetPassword(Phone, Password)
	if err != nil {
		logrus.Errorf("[service.UserResetPassword] %v", err)
		return false
	}
	return true
}

func UserUpdatePhone(token string, Phone int64) bool {
	err := dao.UserUpdatePhone(token, Phone)
	if err != nil {
		logrus.Errorf("[service.UserUpdatePhone] %v", err)
		return false
	}
	return true
}
func UserUpdatePassword(token string, OldPassword string, NewPassword string) bool {
	claim, _ := util.ParseToken(token)
	err := dao.UserUpdatePassword(claim.Phone, OldPassword, NewPassword)
	if err != nil {
		logrus.Errorf("[service.UserUpdatePassword] %v", err)
		return false
	}
	return true
}

func UserGetProfile(Phone int64) (*model.User, error) {
	user, err := dao.GetUserByPhone(Phone)
	if err != nil {
		logrus.Errorf("[service.UserGetProfile] %v", err)
		return nil, err
	}
	return user, nil
}
