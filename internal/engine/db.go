package engine

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"jcz-backend/config"
	"log"
	"os"
	"time"
)

var mysqlCli *gorm.DB
var redisCli *redis.Client

func init() {
	var err error
	//连接MySQL
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetConfig().Mysql.User,
		config.GetConfig().Mysql.Password,
		config.GetConfig().Mysql.Addr,
		config.GetConfig().Mysql.Database,
	)
	gormLog := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})
	mysqlCli, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},

		Logger: gormLog,
	})
	if err != nil {
		logrus.Panicf("[dao.init] database connect %v", err)
	}
	logrus.Infof("[engine] Connected to mysql")
	//连接Redis
	redisCli = redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().Redis.Addr,
		Password: config.GetConfig().Redis.Password,
		DB:       0,
	})

	_, err = redisCli.Ping().Result()
	if err != nil {
		logrus.Panicf("[engine] database connect %v", err)
	}
	logrus.Infof("[engine] Connected to redis")
}

func GetSqlCli() *gorm.DB {
	return mysqlCli
}

func GetRedisCli() *redis.Client {
	return redisCli
}
