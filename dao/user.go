package dao

import (
	"LipLanguage/model"
	"errors"
	"github.com/sirupsen/logrus"
)

func Register(user *model.User) error {
	return DB.Create(&user).Error
}

func GetUserByPhone(phone int64) (*model.User, error) {
	user := &model.User{}
	err := DB.Model(model.User{}).Where("phone = ?", phone).Take(user).Error
	return user, err
}

func GetUserByNickname(nickname string) (*model.User, error) {
	user := &model.User{}
	err := DB.Model(model.User{}).Where("nickname = ?", nickname).Take(user).Error
	return user, err
}

func GetUserByID(ID int64) (*model.User, error) {
	user := &model.User{}
	err := DB.Model(model.User{}).Where("id = ?", ID).Take(user).Error
	return user, err
}

func UserInfoUpdate(Phone int64, info *model.UpdateInfoParam) error {
	user := &model.User{}
	err := DB.Model(model.User{}).Where("phone = ?", Phone).Take(user).Error
	if err != nil {
		return err
	}
	user.Nickname = info.Nickname
	user.Name = info.Name
	user.Email = info.Email
	user.BirthDay = info.Birthday
	user.Gender = info.Gender
	user.HearingDevice = info.HearingDevice
	return DB.Model(model.User{}).Where("id=?", user.ID).Save(user).Error
}
func UserExists(Phone int64) bool {
	var users []model.User
	DB.Model(model.User{}).Where("phone = ?", Phone).Find(&users)
	return len(users) != 0
}

func UserResetPassword(Phone int64, Password string) error {
	user, err := GetUserByPhone(Phone)
	if err != nil {
		return err
	}

	user.Password = Hash256(Password)
	return DB.Model(model.User{}).Where("phone = ?", Phone).Save(user).Error
}

func UserUpdatePhone(token string, Phone int64) error {
	claim, _ := ParseToken(token)
	user, err := GetUserByPhone(claim.Phone)
	if err != nil {
		logrus.Errorf("[dao.UserUpdatePhone] %v", err)
		return err
	}

	user.Phone = Phone
	err = DB.Model(model.User{}).Where("id = ?", user.ID).Save(user).Error
	if err != nil {
		logrus.Errorf("[dao.UserUpdatePhone] %v", err)
		return err
	}
	return nil
}

func UserUpdatePassword(Phone int64, Old string, New string) error {
	user, err := GetUserByPhone(Phone)
	if err != nil {
		logrus.Errorf("[dao.UserUpdatePassword] %v", err)
		return err
	}
	if user.Password != Hash256(Old) {
		return errors.New("PasswordWrong")
	}
	user.Password = Hash256(New)
	return DB.Model(model.User{}).Where("id = ?", user.ID).Save(user).Error
}
