package model

import "gorm.io/gorm"

// LearnRecord 学习数据
type LearnRecord struct {
	gorm.Model
	UserID  int64
	Result  string
	VideoID int64 `gorm:"index"`
	Right   bool
}

type StandardVideo struct {
	gorm.Model
	ID         int64
	Answer     string
	LearnCount int64
	RightCount int64
}
