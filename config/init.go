package config

import (
	"LipLanguage/common"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// 设置配置文件路径
	viper.AddConfigPath("../src")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	logrus.Infof("[config.init] inited config")

	common.SrcPath = viper.GetString("SrcPath")
	common.AIUrl = viper.GetString("AIUrl")

	common.MysqlIpaddr = viper.GetString("MySQL.IpAddr")
	common.MySqlUsername = viper.GetString("MySQL.Username")
	common.MySqlDatabase = viper.GetString("MySQL.Database")
	common.MySqlPassword = viper.GetString("MySQL.Password")
	common.MysqlPort = viper.GetString("MySQL.Port")

	common.RedisIpaddr = viper.GetString("Redis.IpAddr")
	common.RedisPassword = viper.GetString("Redis.Password")
	common.RedisPort = viper.GetString("Redis.Port")

	fmt.Printf(`
Loading Configuration:
SrcPath = %v
AIUrl = %v
MySQL:
Database = %v
	Username = %v
	Password = %v
	IpAddr = %v
	Port = %v
Redis:
	IpAddr = %v
	Password = %v
	Port = %v
`, common.SrcPath, common.AIUrl,
		common.MySqlDatabase, common.MySqlUsername, common.MySqlPassword, common.MysqlIpaddr, common.MysqlPort,
		common.RedisIpaddr, common.RedisPassword, common.RedisPort)

}
