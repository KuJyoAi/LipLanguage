package learn

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"gorm.io/gorm"
	"time"
)

// GetTodayStatisticsByUserID 根据用户ID获取用户的学习统计数据, 如果不存在则根据上一个数据创建
func GetTodayStatisticsByUserID(userID int64) (ret model.LearnStatistics, err error) {
	now := time.Now()
	err = dao.DB.
		Where("user_id = ? AND year = ? AND month = ? AND day = ?",
			userID, now.Year(), now.Month(), now.Day()).
		First(&ret).Error

	// 如果不存在则创建
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 获取上一条数据
			lastStatistics, err := GetLastStatisticsByUserID(userID)
			if err != nil {
				return ret, err
			}
			// 创建今日数据
			ret = model.LearnStatistics{
				Year:        now.Year(),
				Month:       int(now.Month()),
				Day:         now.Day(),
				UserID:      userID,
				TodayLearn:  0,
				TodayMaster: 0,
				TotalLearn:  lastStatistics.TotalLearn,
				TodayTime:   0,
				TotalTime:   lastStatistics.TotalTime,
			}
			err = dao.DB.Create(&ret).Error
			if err != nil {
				return ret, err
			}
			return ret, nil
		} else {
			return ret, err
		}
	}

	return ret, nil
}

// GetLastStatisticsByUserID 根据用户ID获取用户的最后一条学习统计数据, 如果不存在则创建
func GetLastStatisticsByUserID(userID int64) (ret model.LearnStatistics, err error) {
	err = dao.DB.
		Where("user_id = ?", userID).
		Order("create_at desc").
		First(&ret).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 创建第一条统计数据
			ret = model.LearnStatistics{
				Year:   time.Now().Year(),
				Month:  int(time.Now().Month()),
				Day:    time.Now().Day() - 1,
				UserID: userID,
			}
			err = dao.DB.Create(&ret).Error
			if err != nil {
				return ret, err
			}
			return ret, nil
		} else {
			return ret, err
		}
	}
	return ret, nil
}

func GetMonthStatisticsByUserID(userID int64, year int, month int) (ret []model.LearnStatistics, err error) {
	err = dao.DB.
		Where("user_id = ? AND year = ? AND month = ?", userID, year, month).
		Find(&ret).Error
	return ret, err
}
