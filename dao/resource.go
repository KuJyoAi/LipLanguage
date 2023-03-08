package dao

import (
	"LipLanguage/common"
	"LipLanguage/model"
	"gorm.io/gorm"
)

func GetResource(SrcID uint) (*model.Resource, error) {
	ret := &model.Resource{}
	err := DB.Model(model.Resource{}).Where("id=?", SrcID).Take(ret).Error
	return ret, err
}

func CreateResource(path string, Type int64, Filename string) (model.Resource, error) {
	src := model.Resource{
		Model:    gorm.Model{},
		Path:     path,
		Type:     Type,
		Filename: Filename,
	}
	err := DB.Model(model.Resource{}).Create(&src).Error
	return src, err
}

func GetLastVideo() (model.Resource, error) {
	res := model.Resource{}
	err := DB.Model(model.Resource{}).
		Where("type = ?", common.VideoResource).Order("id desc").Take(&res).Error
	return res, err
}
