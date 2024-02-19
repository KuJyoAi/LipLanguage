package model

import (
	"gorm.io/gorm"
)

// LearnRecord 学习记录
type LearnRecord struct {
	gorm.Model

	UserID  int64  `gorm:"index" json:"user_id,omitempty"`
	VideoID int64  `gorm:"video_id,index" json:"video_id,omitempty"`
	Result  string `gorm:"result" json:"result,omitempty"`
	Right   bool   `gorm:"right" json:"right,omitempty"`
	SrcID   string `gorm:"src_id" json:"video_src,omitempty"`
	LipID   string `gorm:"lip_id" json:"video_lip,omitempty"`
}

// UserLearnTime 用户学习时长的记录(每日)
type UserLearnTime struct {
	gorm.Model
	UserID    uint  `gorm:"user_id,uniqueIndex:idx" json:"-"`
	TimeInt   int   `gorm:"time_int,uniqueIndex:idx" json:"time_int"`
	LearnTime int64 `gorm:"learn_time" json:"learn_time"`
}
