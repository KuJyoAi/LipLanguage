package main

import (
	"LipLanguage/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
)

func main() {
	r := gin.Default()
	api.Router(r)
	r.LoadHTMLFiles("html/index.html")
	r.Run()
}

// 初始化日志
func initFile() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	logFile := viper.GetString("LogFile")
	f, err := os.Create(logFile)
	if err != nil {
		panic(err)
	}
	// 使用gin的debug写入
	gin.DefaultWriter = io.MultiWriter(f)
}
