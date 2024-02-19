package service

import "jcz-backend/dao"

func GetResource(SrcID string) ([]byte, error) {
	return dao.GetResourceDataBySrcID(SrcID)
}
