package dao

import "LipLanguage/model"

func AddRouterCounter(record model.RouterCounter) error {
	return DB.Model(model.RouterCounter{}).Create(&record).Error
}
