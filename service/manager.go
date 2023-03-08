package service

import (
	"LipLanguage/common"
	"LipLanguage/dao"
	"LipLanguage/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mime/multipart"
)

func UploadStandardVideo(
	ctx *gin.Context, video *multipart.FileHeader, lip *multipart.FileHeader, answer string) (int64, int64, error) {
	// 写入视频
	// 读取最后一个视频
	LastVideo, err := dao.GetLastVideo()
	if err != nil {
		logrus.Errorf("[service.UploadStandardVideo] %v", err)
		return 0, 0, err
	}
	logrus.Infof("[service.UploadStandardVideo] Read LastVideo %v", LastVideo)

	//创建文件
	videoPath := fmt.Sprintf(common.SrcPath+"/src/standard/%v.mp4", LastVideo.ID+1)
	lipPath := fmt.Sprintf(common.SrcPath+"/src/standard/%v_lip.mp4", LastVideo.ID+1)
	err1 := ctx.SaveUploadedFile(video, videoPath)
	err2 := ctx.SaveUploadedFile(lip, lipPath)
	if err1 != nil || err2 != nil {
		logrus.Errorf("[service.UploadStandardVideo] err1=%v err2=%v \n VideoPath=%v \n LipPath=%v",
			err1, err2, videoPath, lipPath)
		return 0, 0, err1
	}
	logrus.Infof("[service.UploadStandardVideo] Created File: \n VideoPath=%v \n LipPath=%v",
		videoPath, lipPath)

	//资源进入数据库
	VideoResource, err1 := dao.CreateResource(videoPath, common.VideoResource,
		fmt.Sprintf("%v.mp4", LastVideo.ID+1))
	LipResource, err2 := dao.CreateResource(lipPath, common.LipVideoResource,
		fmt.Sprintf("%v_lip.mp4", LastVideo.ID+1))
	if err1 != nil || err2 != nil {
		logrus.Errorf("[service.UploadStandardVideo] err1=%v err2=%v", err1, err2)
		return 0, 0, err1
	}
	logrus.Infof("[service.UploadStandardVideo] Resource Database Saved, StandardVideo ID = %v",
		LastVideo.ID+1)

	// 创建标准视频
	Standard := model.StandardVideo{
		Model:      gorm.Model{},
		Answer:     answer,
		SrcID:      int64(VideoResource.ID),
		LipID:      int64(LipResource.ID),
		LearnCount: 0,
		RightCount: 0,
	}
	err = dao.CreateStandardVideo(Standard)
	if err != nil {
		logrus.Errorf("[service.UploadStandardVideo] %v", err)
		return 0, 0, err
	}
	logrus.Infof("[service.UploadStandardVideo] Standard Saved, StandardVideo ID = %v",
		LastVideo.ID+1)
	return Standard.SrcID, Standard.LipID, nil
}
