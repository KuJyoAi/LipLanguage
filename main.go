package main

import (
	"LipLanguage/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.Router(r)
	r.Run(":8081")
}
