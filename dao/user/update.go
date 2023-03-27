package user

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"errors"
	"github.com/sirupsen/logrus"
	"time"
)

func ResetPassword(Phone int64, Password string) error {
	user, err := GetByPhone(Phone)
	if err != nil {
		return err
	}

	user.Password = dao.Hash256(Password)
	return dao.DB.Model(model.User{}).Where("phone = ?", Phone).Save(user).Error
}

func UpdatePhone(token string, Phone int64) error {
	claim, _ := dao.ParseToken(token)
	user, err := GetByPhone(claim.Phone)
	if err != nil {
		logrus.Errorf("[dao.UpdatePhone] %v", err)
		return err
	}

	user.Phone = Phone
	err = dao.DB.Model(model.User{}).Where("id = ?", user.ID).Save(user).Error
	if err != nil {
		logrus.Errorf("[dao.UpdatePhone] %v", err)
		return err
	}
	return nil
}

func UpdatePassword(Phone int64, Old string, New string) error {
	user, err := GetByPhone(Phone)
	if err != nil {
		logrus.Errorf("[dao.UpdatePassword] %v", err)
		return err
	}
	if user.Password != dao.Hash256(Old) {
		return errors.New("PasswordWrong")
	}
	user.Password = dao.Hash256(New)
	return dao.DB.Model(model.User{}).Where("id = ?", user.ID).Save(user).Error
}

func UpdateInfo(Phone int64, info *model.UpdateInfoParam) error {
	user := &model.User{}
	err := dao.DB.Model(model.User{}).Where("phone = ?", Phone).Take(user).Error
	if err != nil {
		return err
	}
	user.Nickname = info.Nickname
	user.Name = info.Name
	user.Email = info.Email
	user.Gender = info.Gender
	user.HearingDevice = info.HearingDevice
	//转换时间 YYYY-MM-DD
	user.BirthDay, err = time.Parse("2006-01-02", info.Birthday)
	return dao.DB.Model(model.User{}).Where("id=?", user.ID).Save(user).Error
}
