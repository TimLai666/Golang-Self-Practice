package main

import (
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// 允許來自所有來源的 CORS 請求
	// router.Use(cors.Default()) // 使用cors中介軟體

	// 設定路由
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello GO!")
	})

	// 監聽埠
	router.Run("127.0.0.1:9090")
}
