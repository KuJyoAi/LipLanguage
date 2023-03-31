package learn

import (
	"LipLanguage/common"
	"LipLanguage/dao/learn"
	"LipLanguage/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// SaveVideoFile 保存视频 ID为视频结果记录的ID
func SaveVideoFile(record model.LearnRecord, data []byte) error {
	now := time.Now()
	filename := fmt.Sprintf("%d_Video%d_User%d_Time%d_%d_%d.mp4",
		record.ID, record.VideoID, record.UserID, now.Year(), now.Month(), now.Day())
	f, err := os.OpenFile(common.SrcPath+"/src/user/"+filename, os.O_CREATE, 0777)
	defer f.Close()
	if err != nil {
		logrus.Errorf("[service] SaveVideoFile %v", err)
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		logrus.Errorf("[service] SaveVideoFile %v", err)
		return err
	}
	return nil
}

// GetAllStandardVideos 获取所有标准视频, 加上统计信息
func GetAllStandardVideos(UserID int64, limit, offset int, order string) ([]model.StandardVideoResponse, error) {
	data, err := learn.GetAllStandardVideos(UserID, limit, offset, order)
	if err != nil {
		logrus.Errorf("[service.GetAllStandardVideos] %v", err)
		return nil, err
	}
	return data, err
}
