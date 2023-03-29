package model

import (
	"gorm.io/gorm"
	"time"
)

// StandardVideo 标准视频
type StandardVideo struct {
	ID        int64          `gorm:"primaryKey" json:"id,omitempty"`
	Answer    string         `gorm:"answer" json:"answer"`
	SrcID     string         `gorm:"src_id" json:"src_id"`
	LipID     string         `gorm:"lip_id" json:"lip_id"`
	CreateAt  time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
