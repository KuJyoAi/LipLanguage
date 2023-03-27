package user

import (
	"LipLanguage/dao"
	"LipLanguage/model"
)

func Create(user *model.User) error {
	return dao.DB.Create(&user).Error
}

func Exists(Phone int64) bool {
	var users []model.User
	dao.DB.Model(model.User{}).Where("phone = ?", Phone).Find(&users)
	return len(users) != 0
}
