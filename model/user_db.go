package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	AvatarUrl     string    `gorm:"avatar_url"`
	Phone         int64     `gorm:"phone,unique"`
	Email         string    `gorm:"email,unique"`
	Nickname      string    `gorm:"nickname"`
	Name          string    `gorm:"name"`
	Password      string    `gorm:"password"`
	HearingDevice bool      `gorm:"hearing_device"`
	Gender        int       `gorm:"gender"`
	BirthDay      time.Time `gorm:"birthday"`
}
