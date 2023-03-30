package learn

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

// CreateLearnRecord 创建学习记录, 并更新用户的统计数据(表同步更新)
func CreateLearnRecord(record model.LearnRecord) error {
	tx := dao.DB.Begin()
	err := tx.Create(&record).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新用户的统计数据
	err = SyncBasedOnLearnRecord(record, tx)
	if err != nil {
		logrus.Errorf("[dao] SyncBasedOnLearnRecord %v", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// GetLastLearnRecordByUserID 根据用户ID获取用户的最后一条学习记录, 如果不存在则返回错误
func GetLastLearnRecordByUserID(userID int64) (model.LearnRecord, error) {
	lastRecord := model.LearnRecord{}
	err := dao.DB.
		Where("user_id = ?", userID).
		Order("created_at desc").
		First(&lastRecord).Error
	return lastRecord, err
}

// SyncBasedOnLearnRecord 根据学习记录更新用户的统计数据
func SyncBasedOnLearnRecord(record model.LearnRecord, tx *gorm.DB) error {
	// 1. Statistics表
	// (1) 今日学习次数+1, 总学习次数+1
	todayStatistics, err := GetTodayStatisticsByUserID(record.UserID)
	todayStatistics.TodayLearn += 1
	todayStatistics.TotalLearn += 1
	if record.Right {
		todayStatistics.TodayMaster += 1
	}
	// (2) 今日学习时长, 总学习时长
	lastRecord, err := GetLastLearnRecordByUserID(record.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("[dao.SyncBasedOnLearnRecord] GetLastLearnRecordByUserID %v", err)
		return err
	}
	if err == gorm.ErrRecordNotFound {
		// 如果没有上一条记录, 则今日学习时长为1
		todayStatistics.TodayTime = 1
		todayStatistics.TotalTime = 1
	} else {
		// 如果有上一条记录, 则今日学习时长为本次学习时间减去上一次学习时间, 单位为min
		// 10min外的学习记录不计入今日学习时长
		// 不足1min的学习记录按1min计算
		if record.CreatedAt.Sub(lastRecord.CreatedAt) <= 10*time.Minute {
			learnSeconds := int(record.CreatedAt.Sub(lastRecord.CreatedAt).Seconds())
			todayStatistics.TodayTime += (learnSeconds / 60) + 1
			todayStatistics.TotalTime += (learnSeconds / 60) + 1
		}
	}
	err = tx.Save(&todayStatistics).Error
	if err != nil {
		logrus.Errorf("[dao.SyncBasedOnLearnRecord] Statistics Sync %v", err)
		return err
	}

	// 2. StandardVideoCount表
	err = tx.Model(&model.StandardVideoCount{}).
		Where("user_id = ? and video_id = ?", record.UserID, record.VideoID).
		FirstOrCreate(&model.StandardVideoCount{
			UserID:     record.UserID,
			VideoID:    record.VideoID,
			LearnCount: 0,
			LearnTime:  0,
		}).Error
	if err != nil {
		logrus.Errorf("[dao.SyncBasedOnLearnRecord] StandardVideoCount Take %v", err)
		return err
	}
	logrus.Infof("[dao.SyncBasedOnLearnRecord] StandardVideoCount Take %v", err)
	err = tx.Model(&model.StandardVideoCount{}).
		Where("user_id = ? and video_id = ?", record.UserID, record.VideoID).
		Update("learn_count", gorm.Expr("learn_count + ?", 1)).
		Update("learn_time",
			gorm.Expr("learn_time + ?", record.CreatedAt.Sub(lastRecord.CreatedAt).Minutes()+1)).
		Error
	if err != nil {
		logrus.Errorf("[dao.SyncBasedOnLearnRecord] StandardVideoCount Update %v", err)
	}
	return nil
}

func GetStandardVideoLearnRecord(
	UserID int64, VideoID int64, limit int, offset int, order string) (
	data []model.LearnRecordResponse, err error) {
	if order == "" {
		order = "created_at desc"
	}
	err = dao.DB.
		Model(&model.LearnRecord{}).
		Select(`learn_record.src_id
					  learn_record.lip_id
					  learn_record.created_at
					  learn_record.result
					  learn_record.right`).
		Where("user_id = ? and video_id = ?", UserID, VideoID).
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&data).Error
	return
}
