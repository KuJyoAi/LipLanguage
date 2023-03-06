package common

import "time"

const (
	JwtExpireTime = 7 * 24 * time.Hour
	JwtKey        = "HASH256123456"
)

const (
	HttpExpireTime = 30 * time.Second
)

const (
	ManagerAuth = "lip_manager"
)

// Mysql

var MySqlDatabase = ""
var MySqlUsername = ""
var MySqlPassword = ""
var MysqlIpaddr = ""
var MysqlPort = ""

// Redis

var RedisPassword = ""
var RedisIpaddr = ""
var RedisPort = ""

//Others

var SrcPath = ""
var AIUrl = ""
