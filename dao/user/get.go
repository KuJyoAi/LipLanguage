package user

import (
	"LipLanguage/dao"
	"LipLanguage/model"
)

func GetByPhone(phone int64) (*model.User, error) {
	user := &model.User{}
	err := dao.DB.Model(model.User{}).Where("phone = ?", phone).Take(user).Error
	return user, err
}

func GetByNickname(nickname string) (*model.User, error) {
	user := &model.User{}
	err := dao.DB.Model(model.User{}).Where("nickname = ?", nickname).Take(user).Error
	return user, err
}

func GetByID(ID int64) (*model.User, error) {
	user := &model.User{}
	err := dao.DB.Model(model.User{}).Where("id = ?", ID).Take(user).Error
	return user, err
}
