package common

import "time"

const (
	JwtExpireTime = 3 * 24 * time.Hour
	JwtKey        = "HASH256123456"
)

const (
	HttpExpireTime = 30 * time.Second
	AIUrl          = ""
)

const (
	StandardVideoPath = "src"
	TrainVideoPath    = ""
)
