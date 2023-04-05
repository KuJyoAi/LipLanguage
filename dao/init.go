package dao

import (
	"LipLanguage/common"
	_ "LipLanguage/config"
	"LipLanguage/model"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

func init() {
	var err error
	//连接MySQL
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		common.MySqlUsername, common.MySqlPassword,
		common.MysqlIpaddr, common.MysqlPort, common.MySqlDatabase)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panicf("[dao.init] database connect %v", err)
	}
	logrus.Infof("[dao.init] Connected to mysql")
	DB.AutoMigrate(&model.User{},
		&model.LearnStatistics{},
		&model.LearnRecord{},
		&model.StandardVideo{},
		&model.Resource{},
		&model.RouterCounter{},
		&model.StandardVideoCount{},
		&model.Notice{})

	//连接Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", common.RedisIpaddr, common.RedisPort),
		Password: common.RedisPassword,
		DB:       0,
	})
	_, err = RDB.Ping().Result()
	if err != nil {
		logrus.Panicf("[dao.init] database connect %v", err)
	}
	logrus.Infof("[dao.init] Connected to redis")

}
