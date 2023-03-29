package service

import "LipLanguage/dao"

func GetResource(SrcID string) ([]byte, error) {
	return dao.GetResourceDataBySrcID(SrcID)
}
