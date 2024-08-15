package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Testing(c *gin.Context) {

	fmt.Println("Hello world")
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}
