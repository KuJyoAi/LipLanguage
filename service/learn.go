package service

import (
	"LipLanguage/common"
	"LipLanguage/dao"
	"LipLanguage/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"sync"
	"time"
)

func UploadVideo(phone int64, VideoID int64, data *[]byte) (*model.AiPostResponse, error) {
	user, err := dao.GetUserByPhone(phone)
	if err != nil {
		logrus.Errorf("[service] UpdateVideo %v", err)
		return nil, err
	}
	rec := model.LearnRecord{
		Model:   gorm.Model{},
		UserID:  int64(user.ID),
		Result:  "",
		VideoID: VideoID,
		Right:   false,
	}

	// 保存文件到本地
	path, err := SaveTrainVideo(user, VideoID, data)
	if err != nil {
		logrus.Errorf("[service.UploadVideo]%v", err)
		return nil, err
	}

	// 发送给算法
	resp, err := dao.PostVideoPath(path)
	if err != nil {
		logrus.Errorf("[service.UploadVideo]%v", err)
		return nil, err
	}

	// 保存记录, 并传回前端
	wt := sync.WaitGroup{}
	wt.Add(2)

	go func() {
		// 验证是否正确
		rec.Result = resp.Result
		stand, err := dao.GetStandardVideo(VideoID)
		if err != nil {
			logrus.Errorf("[service] UpdateVideo %v", err)
		}
		rec.Right = stand.Answer == resp.Result
		// 更新学习记录
		if rec.Right {
			dao.AddMasterCount(user.ID, 1)
		} else {
			dao.AddLearnCount(user.ID, 1)
		}
		// 存库
		err = dao.SaveLearnRecord(rec)
		if err != nil {
			logrus.Errorf("[service] UpdateVideo %v", err)
		}
		wt.Done()
	}()

	go func() {
		// 写文件
		now := time.Now()
		path = fmt.Sprintf(common.SrcPath+"/src/user/u%v_v%v_%v-%v-%v-%v-%v-%v.mp4",
			user.ID, VideoID, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		err = os.WriteFile(path, *resp.Data, 0777)
		if err != nil {
			logrus.Errorf("[service] UpdateVideo %v", err)
		}
		wt.Done()
	}()

	wt.Wait()
	return &resp, nil
}

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

func GetVideoHistory(VideoID int64, Offset int, Limit int) ([]model.LearnRecord, error) {
	_, err := dao.GetStandardVideo(VideoID)
	if err != nil {
		logrus.Errorf("[service.GetVideoHistory] %v", err)
		return nil, err
	}
	ret, err := dao.GetVideoLearnData(VideoID, Offset, Limit)
	if err != nil {
		logrus.Errorf("[service.GetVideoHistory] %v", err)
		return nil, err
	}
	return *ret, err
}

func GetTodayLearnData(phone int64) (model.LearnStatistics, error) {
	user, err := dao.GetUserByPhone(phone)
	if err != nil {
		logrus.Errorf("[service.GetTodayLearnData] %v", err)
		return model.LearnStatistics{}, err
	}

	// 获取今日数据
	data, err := dao.GetUserTodayStatistics(user.ID)
	if err != nil {
		logrus.Errorf("[service.GetTodayLearnData] %v", err)
		return model.LearnStatistics{}, err
	}

	logrus.Infof("[service.GetTodayLearnData] data:%v", data)

	return data, nil
}

func GetAllStandardVideos(limit, offset int) ([]model.StandardVideo, error) {
	data, err := dao.GetAllStandardVideos(limit, offset)
	if err != nil {
		logrus.Errorf("[service.GetAllStandardVideos] %v", err)
		return nil, err
	}
	return *data, err
}

func SaveTrainVideo(user *model.User, vid int64, data *[]byte) (string, error) {
	now := time.Now()
	path := fmt.Sprintf(common.SrcPath+"/src/user/u%v_v%v_%v-%v-%v-%v-%v-%v.webm",
		user.ID, vid, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	return path, os.WriteFile(path, *data, 0777)
}

func GetDayHistory(limit int, offset int, UserID int64) ([]model.LearnStatistics, error) {
	ret, err := dao.GetDayHistory(limit, offset, UserID)
	if err != nil {
		logrus.Errorf("[service.GetDayHistory] %v", err)
		return []model.LearnStatistics{}, err
	}
	return *ret, err
}
