package learn

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"sync"
)

var AiLock sync.Mutex

func GetUserLastRouters(UserID uint) (*model.RouterCounter, error) {
	ret := model.RouterCounter{}
	err := dao.DB.Where("user_id=?", UserID).Order(" updated_at desc").Take(&ret).Error
	return &ret, err
}
