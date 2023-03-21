package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ReturnIndex(c *gin.Context) {
	fmt.Println("INDEX")
	c.HTML(200, "html/index.html", gin.H{})
}
