package service

import (
	"LipLanguage/dao"
	"github.com/sirupsen/logrus"
	"os"
)

func UploadStandardVideo(data []byte, answer string) error {
	path, err := dao.SaveStandardVideo(data, answer)
	if err != nil {
		logrus.Errorf("[service.UploadStandardVideo] %v path=%v", err, path)
		return err
	}
	// 写入本地
	err = os.WriteFile(path, data, 0777)
	if err != nil {
		logrus.Errorf("[service.UploadStandardVideo] %v path=%v", err, path)
		return err
	}
	return err
}
