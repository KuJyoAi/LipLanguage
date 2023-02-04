package dao

import (
	"LipLanguage/model"
)

func Register(user *model.User) error {
	return DB.Create(&user).Error
}

func GetUserByPhone(phone int64) (*model.User, error) {
	user := &model.User{}
	err := DB.Model(model.User{}).Where("phone = ?", phone).Take(user).Error
	return user, err
}

func UserInfoUpdate(UserID int64, info *model.UpdateInfoParam) error {
	user := &model.User{}
	err := DB.Model(model.User{}).Where("id = ?", UserID).Take(user).Error
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
func UserExists(UserID int64) bool {
	var users []model.User
	DB.Model(model.User{}).Where("id = ?", UserID).Find(&users)
	return len(users) != 0
}
