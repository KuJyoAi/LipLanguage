package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            int64          `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	AvatarUrl     string         `gorm:"avatar_url"`
	Phone         int64          `gorm:"phone,unique"`
	Email         string         `gorm:"email,unique"`
	Nickname      string         `gorm:"nickname"`
	Name          string         `gorm:"name"`
	Password      string         `gorm:"password"`
	HearingDevice bool           `gorm:"hearing_device"`
	Gender        int            `gorm:"gender"`
	BirthDay      time.Time      `gorm:"birthday"`
}

// Notice 通知
type Notice struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	UserID    int64          `gorm:"user_id, index" json:"user_id"`
	Status    int            `gorm:"status" json:"status"`
	Title     string         `gorm:"title" json:"title"`
	Content   string         `gorm:"content" json:"content"`
	CreatedAt time.Time      `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
