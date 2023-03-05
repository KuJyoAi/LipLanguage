package dao

import (
	"LipLanguage/common"
	"LipLanguage/model"
	"LipLanguage/util"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

func GetStandardVideo(ID int64) (model.StandardVideo, error) {
	ret := model.StandardVideo{}
	err := DB.Model(model.StandardVideo{}).Where("id=?", ID).Take(&ret).Error
	return ret, err
}

func SaveLearnRecord(data model.LearnRecord) error {
	return DB.Model(model.LearnRecord{}).Create(&data).Error
}

func GetVideoLearnData(VideoID int64, Offset int, Limit int) (*[]model.LearnRecord, error) {
	var ret []model.LearnRecord
	err := DB.Model(model.LearnRecord{}).
		Where("video_id=?", VideoID).
		Offset(Offset).
		Limit(Limit).
		Find(&ret).Error
	return &ret, err
}

func GetUserLastRouters(UserID uint) (*model.RouterCounter, error) {
	ret := model.RouterCounter{}
	err := DB.Where("user_id=?", UserID).Order(" updated_at desc").Take(&ret).Error
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
	err := DB.Model(model.LearnStatistics{}).Create(&data).Error
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
	err := DB.Model(model.LearnStatistics{}).
		Where("user_id = ?", UserID).
		Order("created_at desc").
		Take(&Statistic).Error
	fmt.Printf("[dao.GetUserTodayStatistics] 查最新的学习数据%+v\n", Statistic)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 用户的第一个数据, 创建
			fmt.Println("创建第一个数据")
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
			err = DB.Model(model.LearnStatistics{}).Save(&Statistic).Error
		} else {
			// 未找到之外的其他错误
			logrus.Errorf("[dao.GetUserTodayStatistics] %v", err)
			return model.LearnStatistics{}, err
		}
	} else {
		// 查到上一次数据, 需要先比较是否为今天的, 不是则创建
		if !util.SameDay(Statistic.Today, time.Now()) {
			Statistic, err = CreateTodayStatistics(UserID, Statistic)
			fmt.Printf("[dao.GetUserTodayStatistics]不是同一天, 创建\n")
			if err != nil {
				return Statistic, err
			}
		}
	}

	// 更新时间
	Statistic, err = UpdateStatisticTime(Statistic, UserID)

	return Statistic, err
}

// AddLearnCount 增加今日学习新词
func AddLearnCount(UserID uint, add int) error {
	Statistic, err := GetUserTodayStatistics(UserID)
	if err != nil {
		return err
	}

	Statistic.TodayLearn += add
	Statistic.TotalLearn += add
	return DB.Model(model.LearnStatistics{}).
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
	return DB.Model(model.LearnStatistics{}).
		Where("id=?", Statistic.ID).
		Save(&Statistic).Error
}

// UpdateStatisticTime 更新时间, 也会保存数据库
// 传入的Statistics必须是最新的
func UpdateStatisticTime(statistics model.LearnStatistics, UserID uint) (model.LearnStatistics, error) {
	// 根据上一次的id寻找之后的统计记录
	var counter []model.RouterCounter
	err := DB.Model(model.RouterCounter{}).
		Where("user_id = ? and id >= ?", UserID, statistics.LastRouterID).
		Find(&counter).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return model.LearnStatistics{}, err
	}

	fmt.Printf("[dao.UpdateStatisticTime] 寻找上一次的统计记录: %+v", counter)

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
	err = DB.Model(model.LearnStatistics{}).Where("id=?", statistics.ID).Save(statistics).Error
	return statistics, err
}

func GetAllStandardVideos(limit int, offset int) (*[]model.StandardVideo, error) {
	var ret []model.StandardVideo
	err := DB.Model(model.StandardVideo{}).
		Offset(offset).Limit(limit).
		Find(&ret).Error
	return &ret, err
}

// PostVideoPath 把视频文件post过去, 发送路径
func PostVideoPath(path string) (model.AiPostResponse, error) {
	// 请求部分
	URL := common.AIUrl + "?VideoPath=" + path
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		logrus.Errorf("[util.PostVideoPath] %v", err)
		return model.AiPostResponse{}, err
	}
	request.Header.Set("Connection", "Keep-Alive")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		logrus.Errorf("[util.PostVideoPath] %v", err)
		return model.AiPostResponse{}, err
	}
	// 读取返回
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("[util.PostVideoPath] %v", err)
	}
	if resp.StatusCode != 200 {
		//有错误
		logrus.Errorf("[util.PostVideoPath] %v:%v",
			resp.StatusCode, resp.Status)
		return model.AiPostResponse{}, errors.New("ai Failed")
	}

	ResLen := data[0]          //结果长度
	video := data[ResLen*3+1:] //使用utf-8编码 长度*3为字节数
	ret := model.AiPostResponse{
		Result: string(data[1 : ResLen*3+1]),
		Data:   &video,
	}
	return ret, nil
}

func GetDayHistory(limit int, offset int, UserID int64) (*[]model.LearnStatistics, error) {
	var ret []model.LearnStatistics
	err := DB.Model(model.LearnStatistics{}).Where("id=?", UserID).
		Offset(offset).Limit(limit).Find(&ret).Error
	return &ret, err
}
