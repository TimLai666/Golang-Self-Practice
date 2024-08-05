package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/post/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": id,
		})
	})

	router.Run(":8000")
}
