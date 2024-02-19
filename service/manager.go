package service

import (
	"github.com/sirupsen/logrus"
	"jcz-backend/dao"
	"jcz-backend/model"
)

func UploadStandardVideo(answer string, video []byte, lip []byte) (Video string, Lip string, err error) {
	videoRes, err := dao.CreateResourceData(video)
	if err != nil {
		logrus.Errorf("[service.UploadStandardVideo]CreateResourceData Video %v", err)
		return
	}
	lipRes, err := dao.CreateResourceData(lip)
	if err != nil {
		logrus.Errorf("[service.UploadStandardVideo]CreateResourceData Lip %v", err)
		return
	}
	standard := model.StandardVideo{
		Answer: answer,
		SrcID:  videoRes.SrcID,
		LipID:  lipRes.SrcID,
	}
	err = dao.CreateStandardVideo(standard)
	if err != nil {
		logrus.Errorf("[service.UploadStandardVideo]CreateStandardVideo %v", err)
		return
	}
	Video = videoRes.SrcID
	Lip = lipRes.SrcID
	return Video, Lip, nil
}
