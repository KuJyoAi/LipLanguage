package learn

import (
	"LipLanguage/dao"
	"LipLanguage/model"
)

func GetStandardVideo(ID int64) (model.StandardVideo, error) {
	ret := model.StandardVideo{}
	err := dao.DB.Model(model.StandardVideo{}).Where("id=?", ID).Take(&ret).Error
	return ret, err
}

func GetAllStandardVideos(limit int, offset int) (*[]model.StandardVideo, error) {
	var ret []model.StandardVideo
	err := dao.DB.Model(model.StandardVideo{}).
		Offset(offset).Limit(limit).
		Find(&ret).Error
	return &ret, err
}
