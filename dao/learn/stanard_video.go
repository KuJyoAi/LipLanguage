package learn

import (
	"LipLanguage/dao"
	"LipLanguage/model"
)

func GetStandardVideo(ID int64) (model.StandardVideo, error) {
	ret := model.StandardVideo{}
	err := dao.DB.Model(model.StandardVideo{}).
		Select(`standard_video.id,
					  standard_video.answer,
					  standard_video.src_id,
					  standard_video.lip_id,	
					  standard_video.created_at,	
					  standard_video_count.learn_count,	
					  standard_video_count.learnt_time`).
		Joins(`LEFT JOIN standard_video_count
					 ON standard_video.id = standard_video_count.video_id`).
		Where("standard_video.id = ?", ID).
		First(&ret).Error
	return ret, err
}

func GetAllStandardVideos(limit int, offset int, Order string) ([]model.StandardVideoResponse, error) {
	var ret []model.StandardVideoResponse
	// 默认按照ID升序
	if Order == "" {
		Order = "id asc"
	}
	err := dao.DB.Model(model.StandardVideo{}).
		Select(`standard_video.id,
		              standard_video.answer,
					  standard_video.src_id,
					  standard_video.lip_id,
					  standard_video.created_at,
					  standard_video_count.learn_count,
 					  standard_video_count.learnt_time`).
		Joins(`LEFT JOIN standard_video_count
                     ON standard_video.id = standard_video_count.video_id`).
		Limit(limit).
		Offset(offset).
		Order(Order).
		Find(&ret).Error
	return ret, err
}
