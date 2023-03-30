package model

// AiPostResponse AI算法传回来的数据
type AiPostResponse struct {
	Result string `json:"result"`
	Data   []byte `json:"data"`
}

// StandardVideoResponse 标准视频的返回值
type StandardVideoResponse struct {
	ID         int64  `json:"id"`
	Answer     string `json:"answer"`
	SrcID      string `json:"src_id"`
	LipID      string `json:"lip_id"`
	CreatedAt  string `json:"created_at"`
	LearnCount int    `json:"learn_count"`
	LearnTime  int    `json:"learn_time"`
}

// StandardVideoLearnRecordResponse 标准视频的学习记录
type StandardVideoLearnRecordResponse struct {
	VideoID int64                 `json:"video_id"`
	Answer  string                `json:"answer"`
	SrcID   string                `json:"src_id"`
	LipID   string                `json:"lip_id"`
	Records []LearnRecordResponse `json:"records"`
}

// LearnRecordResponse 学习记录
type LearnRecordResponse struct {
	SrcID     string `json:"src_id"`
	LipID     string `json:"lip_id"`
	Result    string `gorm:"result" json:"result"`
	Right     bool   `gorm:"right" json:"right"`
	CreatedAt string `gorm:"created_at" json:"created_at"`
}
