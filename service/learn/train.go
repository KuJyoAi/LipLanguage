package learn

import (
	"LipLanguage/dao"
	"LipLanguage/dao/learn"
	"LipLanguage/dao/user"
	"LipLanguage/model"
	"errors"
	"github.com/sirupsen/logrus"
)

// UploadTrainVideo 上传训练视频, 并返回结果
func UploadTrainVideo(phone int64, VideoID int64, data []byte) (ret model.AiPostResponse, err error) {
	User, err := user.GetByPhone(phone)
	if err != nil {
		logrus.Errorf("[service] UpdateVideo %v", err)
		return ret, errors.New("用户不存在")
	}

	// 创建训练视频
	trainVideo, err := dao.CreateResourceData(data)
	if err != nil {
		logrus.Errorf("[service] CreateResourceData%v", err)
		return ret, errors.New("创建训练视频失败")
	}

	// 发送给Ai
	AiRes, err := learn.PostToAi(data)
	if err != nil {
		logrus.Errorf("[service] PostToAi %v", err)
		return ret, err
	}
	// 保存训练视频
	lipVideo, err := dao.CreateResourceData(AiRes.Data)
	if err != nil {
		logrus.Errorf("[service] CreateResourceData%v", err)
		return ret, errors.New("创建训练视频失败")
	}
	// 获取标准视频
	standard, err := learn.GetStandardVideo(VideoID)
	if err != nil {
		logrus.Errorf("[service] GetStandardVideo %v", err)
		return ret, err
	}
	// 保存训练记录
	record := model.LearnRecord{
		UserID:  User.ID,
		VideoID: VideoID,
		Result:  AiRes.Result,
		Right:   AiRes.Result == standard.Answer,
		SrcID:   trainVideo.SrcID,
		LipID:   lipVideo.SrcID,
	}

	err = learn.CreateLearnRecord(record)
	if err != nil {
		logrus.Errorf("[service] CreateLearnRecord %v", err)
		return ret, err
	}

	return AiRes, nil
}
