package learn

import (
	"LipLanguage/dao/learn"
	"LipLanguage/model"
	"github.com/sirupsen/logrus"
)

func GetTodayStatistic(UserID int64) (data model.LearnStatistics, err error) {
	data, err = learn.GetTodayStatisticsByUserID(UserID)
	if err != nil {
		return model.LearnStatistics{}, err
	}
	return data, nil
}

func GetMonthStatistic(UserID int64, year int, month int) (data []model.LearnStatistics, err error) {
	data, err = learn.GetMonthStatisticsByUserID(UserID, year, month)
	if err != nil {
		logrus.Errorf("[service.GetMonthStatistic] GetMonthStatisticsByUserID error: %v", err)
		return nil, err
	}
	return data, nil
}
