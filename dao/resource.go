package dao

import "LipLanguage/model"

func GetResource(SrcID uint) (*model.Resource, error) {
	ret := &model.Resource{}
	err := DB.Model(model.Resource{}).Where("id=?", SrcID).Take(ret).Error
	return ret, err
}
