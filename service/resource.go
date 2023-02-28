package service

import (
	"LipLanguage/dao"
	"github.com/sirupsen/logrus"
	"os"
)

func GetResource(SrcID uint) (*[]byte, error) {
	src, err := dao.GetResource(SrcID)
	if err != nil {
		logrus.Errorf("[api.GetResource]%v", err)
		return nil, err
	}
	data, err := os.ReadFile(src.Path)
	if err != nil {
		logrus.Errorf("[api.GetResource]%v", err)
		return nil, err
	}

	return &data, nil
}
