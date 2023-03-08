package model

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Path     string
	Filename string
	Type     int64 `gorm:"index"`
}
