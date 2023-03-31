package learn

import (
	"LipLanguage/dao"
	"LipLanguage/model"
)

func GetStandardVideo(ID int64) (model.StandardVideo, error) {
	ret := model.StandardVideo{}
	err := dao.DB.Model(model.StandardVideo{}).
		Select(`standard_videos.id,
					  standard_videos.answer,
					  standard_videos.src_id,
					  standard_videos.lip_id,	
					  standard_videos.created_at,	
					  standard_video_counts.learn_count,	
					  standard_video_counts.learn_time`).
		Joins(`LEFT JOIN standard_video_counts
					 ON standard_videos.id = standard_video_counts.video_id`).
		Where("standard_videos.id = ?", ID).
		First(&ret).Error
	return ret, err
}

func GetAllStandardVideos(UserID int64, limit int, offset int, Order string) ([]model.StandardVideoResponse, error) {
	var ret []model.StandardVideoResponse
	// 默认按照ID升序
	if Order == "" {
		Order = "id asc"
	}
	err := dao.DB.Model(model.StandardVideo{}).
		Select(`standard_videos.id,
		              standard_videos.answer,
					  standard_videos.src_id,
					  standard_videos.lip_id,
					  standard_videos.created_at,
					  standard_video_counts.learn_count,
 					  standard_video_counts.learn_time`).
		Joins(`LEFT JOIN standard_video_counts
                     ON standard_videos.id = standard_video_counts.video_id`).
		Where("standard_videos.user_id = ?", UserID).
		Limit(limit).
		Offset(offset).
		Order(Order).
		Find(&ret).Error
	return ret, err
}

func CreateStandardVideo(video model.StandardVideo) error {
	return dao.DB.Create(&video).Error
}
