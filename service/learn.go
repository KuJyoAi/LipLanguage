package service

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"sync"
	"time"
)

func UpdateVideo(phone int64, VideoID int64, data *[]byte) (string, error) {
	user, err := dao.GetUserByPhone(phone)
	if err != nil {
		logrus.Errorf("[service] UpdateVideo %v", err)
		return "", err
	}
	rec := model.LearnRecord{
		Model:   gorm.Model{},
		UserID:  int64(user.ID),
		Result:  "",
		VideoID: VideoID,
		Right:   false,
	}
	//发送给算法和写入文件并行处理获取结果
	wait := sync.WaitGroup{}
	wait.Add(2)
	var err1, err2 error
	go func() {
		//todo 视频裁剪
		rec.Result, err1 = SendVideoToAi(*data)
		wait.Done()
	}()
	go func() {
		err2 = SaveVideoFile(rec, *data)
		wait.Done()
	}()
	wait.Wait()
	if err1 != nil || err2 != nil {
		logrus.Errorf("[service] UpdateVideo %v %v", err1, err2)
		return "", err
	}
	//比对标准结果, 保存到数据库中
	standard, err := dao.GetStandardVideo(VideoID)
	if err != nil {
		logrus.Errorf("[service] UpdateVideo %v", err)
		return "", err
	}
	if standard.Answer == rec.Result {
		rec.Right = true
	}
	//保存记录
	err = dao.SaveLearnRecord(rec)
	if err != nil {
		logrus.Errorf("[service] UpdateVideo %v", err)
		return "", err
	}
	return rec.Result, nil
}

// SaveVideoFile 保存视频 ID为视频结果记录的ID
func SaveVideoFile(record model.LearnRecord, data []byte) error {
	now := time.Now()
	filename := fmt.Sprintf("%d_Video%d_User%d_Time%d_%d_%d.mp4",
		record.ID, record.VideoID, record.UserID, now.Year(), now.Month(), now.Day())
	f, err := os.OpenFile("../src/user/"+filename, os.O_CREATE, 0666)
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

func SendVideoToAi(data []byte) (string, error) {
	return "", nil
}
