package learn

import (
	"LipLanguage/common"
	"LipLanguage/dao/learn"
	"LipLanguage/dao/user"
	"LipLanguage/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
	"sync"
	"time"
)

// UploadVideo 上传训练视频, 并返回结果
func UploadVideo(ctx *gin.Context, phone int64, VideoID int64, data *multipart.FileHeader) (*model.AiPostResponse, error, bool) {
	user, err := user.GetByPhone(phone)
	if err != nil {
		logrus.Errorf("[service] UpdateVideo %v", err)
		return nil, err, false
	}
	rec := model.LearnRecord{
		Model:   gorm.Model{},
		UserID:  int64(user.ID),
		Result:  "",
		VideoID: VideoID,
		Right:   false,
	}

	// 保存文件到本地
	logrus.Infof("Saving File: name=%v, time=%v", data.Filename, time.Now())
	path, err := SaveTrainVideo(ctx, user, VideoID, data)
	if err != nil {
		logrus.Errorf("[service.UploadVideo]%v", err)
		return nil, err, false
	} else {
		logrus.Infof("Saved File:%v time=%v", path, time.Now())
	}

	// 发送给算法
	logrus.Infof("Sending File to AI: path=%v, time=%v", path, time.Now())
	resp, err, ok := learn.PostVideoPath(path)

	if err != nil {
		logrus.Errorf("[service.UploadVideo]%v", err)
		if ok {
			// AI算法出错
			logrus.Errorf("AI Failed = %v time = %v", path, time.Now())
			return nil, err, true
		}
		return nil, err, false
	}
	logrus.Infof("AI Result=%v time=%v, data_length=%v bytes", resp.Result, time.Now(), len(*resp.Data))
	// 保存记录, 并传回前端
	wt := sync.WaitGroup{}
	wt.Add(2)

	go func() {
		// 验证是否正确
		rec.Result = resp.Result
		stand, err := learn.GetStandardVideo(VideoID)
		if err != nil {
			logrus.Errorf("[service] UpdateVideo %v", err)
		}
		rec.Right = stand.Answer == resp.Result
		// 更新学习记录
		if rec.Right {
			learn.AddMasterCount(user.ID, 1)
		} else {
			learn.AddLearnCount(user.ID, 1)
		}
		// 存库
		err = learn.SaveLearnRecord(rec)
		if err != nil {
			logrus.Errorf("[service] UpdateVideo %v", err)
		}
		logrus.Infof("Saved Record time=%v", time.Now())
		wt.Done()
	}()

	go func() {
		// 写文件
		now := time.Now()
		path = fmt.Sprintf(common.SrcPath+"/src/user/u%v_v%v_%v_%v_%v_%v_%v_%v.webm",
			user.ID, VideoID, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		err = os.WriteFile(path, *resp.Data, 0777)
		if err != nil {
			logrus.Errorf("[service] UpdateVideo %v", err)
		}
		logrus.Infof("Saved Lip Video time=%v", time.Now())
		wt.Done()
	}()

	wt.Wait()
	return &resp, nil, true
}

func SaveTrainVideo(ctx *gin.Context, user *model.User, vid int64, data *multipart.FileHeader) (string, error) {
	now := time.Now()
	path := fmt.Sprintf(common.SrcPath+"/src/user/u%v_v%v_%v_%v_%v_%v_%v_%v.webm",
		user.ID, vid, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	return path, ctx.SaveUploadedFile(data, path)
}
