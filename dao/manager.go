package dao

import (
	"jcz-backend/model"
)

func CreateStandardVideo(video model.StandardVideo) error {
	return DB.Model(model.StandardVideo{}).Create(&video).Error
}
