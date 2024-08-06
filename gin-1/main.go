package main

import (
	"gin-1/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// router.POST("/post/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.JSON(200, gin.H{
	// 		"message": id,
	// 	})
	// })

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)
	router.Run(":8000")
}
