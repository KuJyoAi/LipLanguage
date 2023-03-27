package learn

import (
	"LipLanguage/dao"
	"LipLanguage/model"
	"LipLanguage/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

// SaveLearnRecord 保存学习记录
func SaveLearnRecord(data model.LearnRecord) error {
	return dao.DB.Model(model.LearnRecord{}).Create(&data).Error
}

// GetVideoLearnData 获取视频的学习记录
func GetVideoLearnData(VideoID int64, Offset int, Limit int) (*[]model.LearnRecord, error) {
	var ret []model.LearnRecord
	err := dao.DB.Model(model.LearnRecord{}).
		Where("video_id=?", VideoID).
		Offset(Offset).
		Limit(Limit).
		Find(&ret).Error
	return &ret, err
}

// CreateTodayStatistics
// 创建今天的记录, 并返回创建的记录
// 没有检查今天是否有, 所以调用前需要做检查
func CreateTodayStatistics(UserID uint, LastRecord model.LearnStatistics) (model.LearnStatistics, error) {
	// 根据上一天的数据创建今天的
	data := model.LearnStatistics{
		Model:        gorm.Model{},
		UserID:       UserID,
		TodayLearn:   0,
		TodayMaster:  0,
		TotalLearn:   LastRecord.TotalLearn,
		TodayTime:    0,
		TotalTime:    LastRecord.TotalTime,
		LastRouterID: LastRecord.LastRouterID,
		Today:        time.Now(),
	}
	err := dao.DB.Model(model.LearnStatistics{}).Create(&data).Error
	if err != nil {
		logrus.Errorf("[dao.CreateTodayStatistics] %v", err)
	}
	return data, err
}

// GetUserTodayStatistics
// 获取用户今天的学习数据
func GetUserTodayStatistics(UserID uint) (model.LearnStatistics, error) {
	// 查最新的学习数据
	Statistic := model.LearnStatistics{}
	err := dao.DB.Model(model.LearnStatistics{}).
		Where("user_id = ?", UserID).
		Order("created_at desc").
		Take(&Statistic).Error
	//fmt.Printf("[dao.GetUserTodayStatistics] 查最新的学习数据%+v\n", Statistic)
	logrus.Infof("[dao.GetUserTodayStatistics] User: %v, Get Last Learn Data", UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 用户的第一个数据, 创建
			Statistic = model.LearnStatistics{
				Model:        gorm.Model{},
				UserID:       UserID,
				TodayLearn:   0,
				TodayMaster:  0,
				TotalLearn:   0,
				TodayTime:    0,
				TotalTime:    0,
				LastRouterID: 0,
				Today:        time.Now(),
			}
			err = dao.DB.Model(model.LearnStatistics{}).Save(&Statistic).Error
			//fmt.Println("Today's First Data")
			logrus.Infof("[dao.GetUserTodayStatistics] User: %v, Today's First Data", UserID)
		} else {
			// 未找到之外的其他错误
			logrus.Errorf("[dao.GetUserTodayStatistics] %v", err)
			return model.LearnStatistics{}, err
		}
	} else {
		// 查到上一次数据, 需要先比较是否为今天的, 不是则创建
		if !util.SameDay(Statistic.Today, time.Now()) {
			Statistic, err = CreateTodayStatistics(UserID, Statistic)
			//fmt.Printf("[dao.GetUserTodayStatistics]不是同一天, 创建\n")
			logrus.Infof("[dao.GetUserTodayStatistics] User: %v, Not Same Day, Create", UserID)
			if err != nil {
				return Statistic, err
			}
		}
	}

	// 更新时间
	Statistic, err = UpdateStatisticTime(Statistic, UserID)

	return Statistic, err
}

// UpdateStatisticTime 更新时间, 也会保存数据库
// 传入的Statistics必须是最新的
func UpdateStatisticTime(statistics model.LearnStatistics, UserID uint) (model.LearnStatistics, error) {
	// 根据上一次的id寻找之后的统计记录
	var counter []model.RouterCounter
	err := dao.DB.Model(model.RouterCounter{}).
		Where("user_id = ? and id >= ?", UserID, statistics.LastRouterID).
		Find(&counter).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return model.LearnStatistics{}, err
	}

	fmt.Printf("[dao.UpdateStatisticTime] User: %v, Num of LastRecords: %+v",
		UserID, len(counter))

	// 不足两个记录, 不需要更新
	if len(counter) < 2 {
		return statistics, nil
	}

	// 计算时间
	period := 0.0
	for i := 0; i < len(counter)-1; i++ {
		t := counter[i+1].CreatedAt.Sub(counter[i].CreatedAt)
		// 小于10分钟内的操作认为用户学习中
		if t <= 10*time.Minute {
			period += t.Minutes()
		}
		// 更新id
		statistics.LastRouterID = counter[i].ID
	}

	statistics.TodayTime += int(period)
	statistics.TotalTime += int(period)
	// 保存到数据库中
	err = dao.DB.Model(model.LearnStatistics{}).Where("id=?", statistics.ID).Save(statistics).Error
	return statistics, err
}

func GetDayHistory(limit int, offset int, UserID int64) (*[]model.LearnStatistics, error) {
	var ret []model.LearnStatistics
	err := dao.DB.Model(model.LearnStatistics{}).Where("id=?", UserID).
		Offset(offset).Limit(limit).Find(&ret).Error
	return &ret, err
}

// AddLearnCount 增加今日学习新词
func AddLearnCount(UserID uint, add int) error {
	Statistic, err := GetUserTodayStatistics(UserID)
	if err != nil {
		return err
	}

	Statistic.TodayLearn += add
	Statistic.TotalLearn += add
	return dao.DB.Model(model.LearnStatistics{}).
		Where("id=?", Statistic.ID).
		Save(&Statistic).Error
}

// AddMasterCount 增加今日掌握新词
func AddMasterCount(UserID uint, add int) error {
	Statistic, err := GetUserTodayStatistics(UserID)
	if err != nil {
		return err
	}

	Statistic.TotalLearn += add
	Statistic.TodayMaster += add
	Statistic.TotalLearn += add
	return dao.DB.Model(model.LearnStatistics{}).
		Where("id=?", Statistic.ID).
		Save(&Statistic).Error
}
