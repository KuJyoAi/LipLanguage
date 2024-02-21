package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Type   string `gorm:"type" json:"type"`
	Answer string `gorm:"answer" json:"answer"`
	Src    string `gorm:"src" json:"src"`   // 资源文件
	Desc   string `gorm:"desc" json:"desc"` // 描述
}

type UserLearnRecord struct {
	gorm.Model
	UserID     uint   `gorm:"user_id" json:"user_id"`
	QuestionID uint   `gorm:"question_id" json:"question_id"`
	Result     string `gorm:"result" json:"result"`
	Right      bool   `gorm:"right" json:"right"`
}
