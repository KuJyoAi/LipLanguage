package model

import (
	"gorm.io/gorm"
	"time"
)

// LearnRecord 学习数据
type LearnRecord struct {
	gorm.Model `json:"gorm-.-model"`
	UserID     int64  `gorm:"index" json:"user_id,omitempty"`
	Result     string `json:"result,omitempty"`
	VideoID    int64  `gorm:"index" json:"video_id,omitempty"`
	Right      bool   `json:"right,omitempty"`
}

// StandardVideo 标准视频
type StandardVideo struct {
	gorm.Model
	Answer     string `gorm:"answer"`
	Path       string `gorm:"path" json:"-"`
	LearnCount int64  `gorm:"learn_count" json:"-"`
	RightCount int64  `gorm:"right_count" json:"-"`
}

// RouterCounter 统计路由调用
type RouterCounter struct {
	gorm.Model
	UserID int64  `gorm:"user_id,index"`
	Path   string `gorm:"path"`
}

// LearnStatistics 学习数据统计
type LearnStatistics struct {
	gorm.Model
	UserID       uint `gorm:"user_id,index"`
	TodayLearn   int  `gorm:"today_learn"`
	TodayMaster  int  `gorm:"today_master"`
	TotalLearn   int  `gorm:"total_learn"`
	LastRouterID uint `gorm:"last_router_id"`
	TodayTime    int  `gorm:"today_time"`
	TotalTime    int  `gorm:"total_time"`
	Today        time.Time
}

// AiPostResponse AI算法传回来的数据
type AiPostResponse struct {
	Result string
	Data   *[]byte
}
