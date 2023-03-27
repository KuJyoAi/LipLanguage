package learn

import (
	"LipLanguage/dao/learn"
	"LipLanguage/dao/user"
	"LipLanguage/model"
	"github.com/sirupsen/logrus"
)

func GetDayHistory(limit int, offset int, UserID int64) ([]model.LearnStatistics, error) {
	ret, err := learn.GetDayHistory(limit, offset, UserID)
	if err != nil {
		logrus.Errorf("[service.GetDayHistory] %v", err)
		return []model.LearnStatistics{}, err
	}
	return *ret, err
}

func GetVideoHistory(VideoID int64, Offset int, Limit int) ([]model.LearnRecord, error) {
	_, err := learn.GetStandardVideo(VideoID)
	if err != nil {
		logrus.Errorf("[service.GetVideoHistory] %v", err)
		return nil, err
	}
	ret, err := learn.GetVideoLearnData(VideoID, Offset, Limit)
	if err != nil {
		logrus.Errorf("[service.GetVideoHistory] %v", err)
		return nil, err
	}
	return *ret, err
}

func GetTodayLearnData(phone int64) (model.LearnStatistics, error) {
	user, err := user.GetByPhone(phone)
	if err != nil {
		logrus.Errorf("[service.GetTodayLearnData] %v", err)
		return model.LearnStatistics{}, err
	}

	// 获取今日数据
	data, err := learn.GetUserTodayStatistics(user.ID)
	if err != nil {
		logrus.Errorf("[service.GetTodayLearnData] %v", err)
		return model.LearnStatistics{}, err
	}

	logrus.Infof("[service.GetTodayLearnData] USER:%v", data.UserID)

	return data, nil
}
