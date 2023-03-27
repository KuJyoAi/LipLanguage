package service

import (
	"LipLanguage/dao"
	"LipLanguage/dao/user"
	"LipLanguage/model"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func Register(Phone int64, Password string) (string, error) {

	User := model.User{
		Model:         gorm.Model{},
		AvatarUrl:     "",
		Phone:         Phone,
		Email:         "",
		Name:          "",
		Password:      dao.Hash256(Password),
		HearingDevice: false,
		Gender:        0,
		BirthDay:      time.Date(2050, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	err := user.Create(&User)
	if err != nil {
		logrus.Errorf("[service] Create %v", err)
		return "", err
	}

	token, err := dao.GenerateToken(User.Phone, User.Nickname, int64(User.ID))

	if err != nil {
		logrus.Errorf("[service] Create %v", err)
		return "", err
	}

	return token, err
}

func LoginByPhone(Phone int64, Password string) (model.LoginResponse, error) {
	User, err := user.GetByPhone(Phone)
	if err != nil {
		logrus.Errorf("[service.Login] %v", err)
		return model.LoginResponse{}, err
	}

	if dao.Hash256(Password) != User.Password {
		logrus.Errorf("[service.Login] %v", err)
		return model.LoginResponse{}, errors.New("PasswordError")
	}

	token, err := dao.GenerateToken(User.Phone, User.Nickname, int64(User.ID))
	if err != nil {
		logrus.Errorf("[service.Login] %v", err)
		return model.LoginResponse{}, err
	}
	birthDay := User.BirthDay.Format("2006-01-02")
	res := model.LoginResponse{
		Token:         token,
		AvatarUrl:     User.AvatarUrl,
		Phone:         fmt.Sprintf("%d", User.Phone),
		Email:         User.Email,
		Name:          User.Name,
		Nickname:      User.Nickname,
		HearingDevice: User.HearingDevice,
		Gender:        User.Gender,
		BirthDay:      birthDay,
		UserID:        User.ID,
	}
	return res, err
}

func LoginByNickname(Nickname string, Password string) (model.LoginResponse, error) {
	User, err := user.GetByNickname(Nickname)
	if err != nil {
		logrus.Errorf("[service.Login] %v", err)
		return model.LoginResponse{}, err
	}

	if dao.Hash256(Password) != User.Password {
		logrus.Errorf("[service.Login] %v", err)
		return model.LoginResponse{}, errors.New("PasswordError")
	}

	token, err := dao.GenerateToken(User.Phone, User.Nickname, int64(User.ID))
	if err != nil {
		logrus.Errorf("[service.Login] %v", err)
		return model.LoginResponse{}, err
	}
	birthDay := User.BirthDay.Format("2006-01-02")
	res := model.LoginResponse{
		Token:         token,
		AvatarUrl:     User.AvatarUrl,
		Phone:         fmt.Sprintf("%d", User.Phone),
		Email:         User.Email,
		Name:          User.Name,
		Nickname:      User.Nickname,
		HearingDevice: User.HearingDevice,
		Gender:        User.Gender,
		BirthDay:      birthDay,
		UserID:        User.ID,
	}
	return res, err
}

func UserInfoUpdate(token string, info *model.UpdateInfoParam) error {
	claim, err := dao.ParseToken(token)
	if err != nil {
		logrus.Errorf("[service.UpdateInfo]ParseToken %v", err)
		return err
	}

	err = user.UpdateInfo(claim.Phone, info)
	if err != nil {
		logrus.Errorf("[service.UpdateInfo] %v", err)
	}
	return err
}

func UserVerify(Phone int64, Name string, Email string) (string, bool) {
	User, err := user.GetByPhone(Phone)
	if err != nil {
		logrus.Errorf("[service.UserVerify] %v", err)
		return "", false

	}

	if User.Email == Email && User.Name == Name {
		//生成验证码, 5分钟有效
		token, err := dao.GenerateTokenExpires(User.Phone, User.Nickname, int64(User.ID), 5*time.Minute)
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
	err := user.ResetPassword(Phone, Password)
	if err != nil {
		logrus.Errorf("[service.ResetPassword] %v", err)
		return false
	}
	return true
}

func UserUpdatePhone(token string, Phone int64) bool {
	err := user.UpdatePhone(token, Phone)
	if err != nil {
		logrus.Errorf("[service.UpdatePhone] %v", err)
		return false
	}
	return true
}
func UserUpdatePassword(token string, OldPassword string, NewPassword string) bool {
	claim, _ := dao.ParseToken(token)
	err := user.UpdatePassword(claim.Phone, OldPassword, NewPassword)
	if err != nil {
		logrus.Errorf("[service.UpdatePassword] %v", err)
		return false
	}
	return true
}

func UserGetProfile(Phone int64) (*model.User, error) {
	User, err := user.GetByPhone(Phone)
	if err != nil {
		logrus.Errorf("[service.UserGetProfile] %v", err)
		return nil, err
	}
	return User, nil
}
