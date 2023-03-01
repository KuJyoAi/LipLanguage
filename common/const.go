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
	StandardVideoPath = ""
	TrainVideoPath    = ""
)

const (
	ManagerAuth = "lip_manager"
)

const (
	MySqlUsername = ""
	MySqlPassword = ""
	MysqlIpaddr   = ""
	MysqlPort     = ""
)

const (
	// RedisPassword RedisUsername = ""
	RedisPassword = ""
	RedisIpaddr   = ""
	RedisPort     = ""
)