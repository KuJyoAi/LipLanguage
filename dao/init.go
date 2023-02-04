package dao

import (
	"LipLanguage/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/lip?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panicf("[dao.init] database connect %v", err)
	}
	logrus.Infof("[dao.init] Connected to mysql")
	DB.AutoMigrate(&model.User{})
}
