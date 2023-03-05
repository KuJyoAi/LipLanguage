package dao

import (
	"LipLanguage/model"
	"fmt"
	"gorm.io/gorm"
)

func SaveStandardVideo(data []byte, answer string) (string, error) {
	video := model.StandardVideo{
		Model:      gorm.Model{},
		Answer:     answer,
		SrcID:      0,
		LipID:      0,
		LearnCount: 0,
		RightCount: 0,
	}
	err := DB.Model(model.StandardVideo{}).Create(&video).Error
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf(
		"src/standard/%v_%v.mp4",
		video.ID, answer)
	err = DB.Model(model.StandardVideo{}).Where("id=?", video.ID).Save(&video).Error
	return path, err
}
