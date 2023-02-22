package dao

import "LipLanguage/model"

func GetStandardVideo(ID int64) (model.StandardVideo, error) {
	ret := model.StandardVideo{}
	err := DB.Model(model.StandardVideo{}).Where("id=?", ID).Take(&ret).Error
	return ret, err
}

func SaveLearnRecord(data model.LearnRecord) error {
	return DB.Model(model.LearnRecord{}).Create(&data).Error
}
