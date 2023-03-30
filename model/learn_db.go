package model

import (
	"gorm.io/gorm"
	"time"
)

// LearnRecord 学习数据记录
type LearnRecord struct {
	ID        int64          `gorm:"primaryKey" json:"id,omitempty"`
	UserID    int64          `gorm:"index" json:"user_id,omitempty"`
	VideoID   int64          `gorm:"video_id,index" json:"video_id,omitempty"`
	Result    string         `gorm:"result" json:"result,omitempty"`
	Right     bool           `gorm:"right" json:"right,omitempty"`
	SrcID     string         `gorm:"src_id" json:"video_src,omitempty"`
	LipID     string         `gorm:"lip_id" json:"video_lip,omitempty"`
	CreatedAt time.Time      `gorm:"created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at, index" json:"-"`
}

// LearnStatistics 学习数据统计
type LearnStatistics struct {
	ID     int64 `gorm:"primaryKey" json:"id"`
	UserID int64 `gorm:"user_id,index" json:"user_id"`

	// 学习情况
	TodayLearn  int `gorm:"today_learn" json:"today_learn"`
	TodayMaster int `gorm:"today_master" json:"today_master"`
	TotalLearn  int `gorm:"total_learn" json:"total_learn"`
	TodayTime   int `gorm:"today_time" json:"today_time"`
	TotalTime   int `gorm:"total_time" json:"total_time"`

	// 查询条件
	Year      int            `gorm:"year index:idx" json:"-"`
	Month     int            `gorm:"month index:idx" json:"-"`
	Day       int            `gorm:"day" json:"-"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at, index" json:"-"`
}

// StandardVideoCount 用户对每个标准视频的学习次数
type StandardVideoCount struct {
	ID      int64 `gorm:"primaryKey" json:"id,omitempty"`
	UserID  int64 `gorm:"user_id,index:idx" json:"user_id"`
	VideoID int64 `gorm:"video_id,index:idx" json:"video_id"`
	// 统计数据
	LearnCount int `gorm:"count" json:"learn_count"`
	LearnTime  int `gorm:"learn_time" json:"learn_time"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// RouterCounter 统计路由调用
type RouterCounter struct {
	ID        int64          `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt time.Time      `gorm:"created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at, index" json:"-"`
	UserID    int64          `gorm:"user_id,index"`
	Path      string         `gorm:"path"`
}
