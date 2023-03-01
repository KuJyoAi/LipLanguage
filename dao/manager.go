package dao

import (
	"LipLanguage/common"
	"LipLanguage/model"
	"fmt"
	"gorm.io/gorm"
)

func SaveStandardVideo(data []byte, answer string) (string, error) {
	video := model.StandardVideo{
		Model:      gorm.Model{},
		Answer:     answer,
		Path:       "",
		LearnCount: 0,
		RightCount: 0,
	}
	err := DB.Model(model.StandardVideo{}).Create(&video).Error
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf(
		common.StandardVideoPath+"/src/standard/%v_%v.mp4",
		video.ID, answer)
	video.Path = path
	err = DB.Model(model.StandardVideo{}).Where("id=?", video.ID).Save(&video).Error
	return path, err
}
