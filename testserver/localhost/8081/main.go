package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.GET("/", func(c *gin.Context) {
		time.Sleep(4 * time.Second)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.Run(":8081")
}
