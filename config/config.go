package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
)

type Config struct {
	ListenPort  string `json:"listen_port" mapstructure:"listen_port"`
	AIUrl       string `json:"ai_url" mapstructure:"ai_url"`
	StoragePath string `json:"storage_path" mapstructure:"storage_path"` // 存储文件路径

	Mysql struct {
		Addr     string `json:"addr" mapstructure:"addr"`
		User     string `json:"user" mapstructure:"user"`
		Password string `json:"password" mapstructure:"password"`
		Database string `json:"database" mapstructure:"database"`
	} `json:"mysql" mapstructure:"mysql"`

	Redis struct {
		Addr     string `json:"addr" mapstructure:"addr"`
		Password string `json:"password" mapstructure:"password"`
	} `json:"redis" mapstructure:"redis"`
}

var c Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}

		if c.StoragePath == "" {
			c.StoragePath = "./data"
		}
		if _, err := os.Stat(c.StoragePath); err != nil {
			if err := os.MkdirAll(c.StoragePath, os.ModePerm); err != nil {
				panic(err)
			}
		}

		if c.ListenPort == "" {
			c.ListenPort = "8080"
		}

		if c.AIUrl == "" {
			logrus.Warn("AIUrl is empty!")
		}
	})
	return &c
}
