package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AvatarUrl     string `gorm:"avatar_url"`
	Phone         int64  `gorm:"phone,unique"`
	Email         string `gorm:"email,unique"`
	Nickname      string `gorm:"nickname"`
	Name          string `gorm:"name"` // 真实姓名
	Password      string `gorm:"password" json:"-"`
	HearingLevel  int    `gorm:"hearing_level"` // 听力等级
	HearingDevice bool   `gorm:"hearing_device"`
	Gender        string `gorm:"gender"`
	BirthDay      string `gorm:"birthday"`
}

// Notice 通知
type Notice struct {
	gorm.Model

	UserID  int64  `gorm:"user_id,index" json:"-"`
	Read    bool   `gorm:"read,default:false" json:"read"`
	Title   string `gorm:"title" json:"title"`
	Content string `gorm:"content" json:"content"`
}
