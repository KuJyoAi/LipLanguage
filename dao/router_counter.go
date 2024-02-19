package dao

import "jcz-backend/model"

func AddRouterCounter(record model.RouterCounter) error {
	return DB.Model(model.RouterCounter{}).Create(&record).Error
}
