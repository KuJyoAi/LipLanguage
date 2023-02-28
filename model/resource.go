package model

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Path string
}
