package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ReturnIndex(c *gin.Context) {
	fmt.Println("INDEX")
	c.HTML(200, "index.html", gin.H{})
}
