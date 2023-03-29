package model

import (
	"gorm.io/gorm"
	"time"
)

type Resource struct {
	ID        int64          `gorm:"primaryKey" json:"id,omitempty"`
	SrcID     string         `gorm:"src_id, index" json:"src_id"`
	Path      string         `gorm:"path"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
