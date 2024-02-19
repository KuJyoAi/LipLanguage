package utils

import (
	"fmt"
	"io"
	"jcz-backend/model"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	dsn := fmt.Sprintf(
		"root:deepkw_lipread@tcp(127.0.0.1:3306)/lipread?charset=utf8mb4&parseTime=True&loc=Local",
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 获取资源
	var records []model.LearnRecord
	err = DB.Model(&model.LearnRecord{}).
		Where("user_id = ?", 9).
		Find(&records).Error
	if err != nil {
		panic(err)
	}
	for _, r := range records {
		res := model.Resource{}
		DB.Model(model.Resource{}).
			Where("src_id = ?", r.SrcID).First(&res)
		logrus.Infof("src_id: %v, path: %v", r.SrcID, res.Path)
		f, _ := os.Open(res.Path)

		data, _ := io.ReadAll(f)
		target, _ := os.Create(fmt.Sprintf("/root/src/tmp/%v_%v.webm", r.Result, r.SrcID))
		_, err := target.Write(data)

		if err != nil {
			logrus.Errorf("write file error: %v", err)
		}
	}
}
