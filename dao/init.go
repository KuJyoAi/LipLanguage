package dao

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"jcz-backend/internal/engine"
)

var DB *gorm.DB
var RDB *redis.Client

func init() {
	DB = engine.GetSqlCli()
	RDB = engine.GetRedisCli()
}
