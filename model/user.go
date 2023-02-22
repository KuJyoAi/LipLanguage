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
	Nickname      string    `gorm:"nickname,unique"`
	Name          string    `gorm:"name,unique"`
	Password      string    `gorm:"password"`
	HearingDevice bool      `gorm:"hearing_device"`
	Gender        int       `gorm:"gender"`
	BirthDay      time.Time `gorm:"birthday"`
}

type RegisterParam struct {
	Phone    int64  `json:"phone"`
	Password string `json:"password"`
}

type LoginParam struct {
	Nickname string `json:"nickname"`
	Phone    int64  `json:"phone"`
	Password string `json:"password,required"`
}

type UpdateInfoParam struct {
	Nickname      string    `json:"nickname"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Birthday      time.Time `json:"birthday"`
	Gender        int       `json:"gender"`
	HearingDevice bool      `json:"hearing_device"`
}

type UserVerifyParam struct {
	Phone int64  `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type ResetPasswordParam struct {
	Phone    int64  `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}
type UpdatePhoneParam struct {
	Phone int64 `json:"phone,omitempty"`
}

type UpdatePasswordParam struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
