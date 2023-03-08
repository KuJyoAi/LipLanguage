package dao

import (
	"LipLanguage/model"
)

func CreateStandardVideo(video model.StandardVideo) error {
	return DB.Model(model.StandardVideo{}).Create(&video).Error
}
