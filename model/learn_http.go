package model

// AiPostResponse AI算法传回来的数据
type AiPostResponse struct {
	Result string
	Data   []byte
}

// StandardVideoResponse 标准视频的返回值
type StandardVideoResponse struct {
	ID         int64  `json:"id,omitempty"`
	Answer     string `json:"answer,omitempty"`
	SrcID      string `json:"src_id,omitempty"`
	LipID      string `json:"lip_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	LearnCount int    `json:"learn_count,omitempty"`
	LearntTime string `json:"learnt_time,omitempty"`
}

// StandardVideoLearnRecordResponse 标准视频的学习记录
type StandardVideoLearnRecordResponse struct {
	VideoID int64                 `json:"video_id,omitempty"`
	Answer  string                `json:"answer,omitempty"`
	SrcID   string                `json:"src_id,omitempty"`
	LipID   string                `json:"lip_id,omitempty"`
	Records []LearnRecordResponse `json:"records,omitempty"`
}

// LearnRecordResponse 学习记录
type LearnRecordResponse struct {
	SrcID     string `json:"src_id,omitempty"`
	LipID     string `json:"lip_id,omitempty"`
	Result    string `gorm:"result" json:"result,omitempty"`
	Right     bool   `gorm:"right" json:"right,omitempty"`
	CreatedAt string `gorm:"created_at" json:"created_at,omitempty"`
}
