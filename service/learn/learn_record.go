package learn

import (
	"jcz-backend/dao/learn"
	"jcz-backend/model"
)

func GetStandardVideoLearnRecord(
	UserID int64, VideoID int64, limit int, offset int, order string) (
	data model.StandardVideoLearnRecordResponse, err error) {

	learnData, err := learn.GetStandardVideoLearnRecord(UserID, VideoID, limit, offset, order)
	if err != nil {
		return data, err
	}

	standardVideo, err := learn.GetStandardVideo(VideoID)
	if err != nil {
		return data, err
	}
	if learnData == nil {
		learnData = make([]model.LearnRecordResponse, 0)
	}
	data = model.StandardVideoLearnRecordResponse{
		VideoID: VideoID,
		Answer:  standardVideo.Answer,
		SrcID:   standardVideo.SrcID,
		LipID:   standardVideo.LipID,
		Records: learnData,
	}

	return data, nil
}
