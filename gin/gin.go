package main

import (
	"golang/gin/app"

	"github.com/gin-gonic/gin"
) // "github.com/gin-contrib/cors"

func main() {
	gin.SetMode(gin.DebugMode)

	// 建立一個 Gin 引擎
	router := gin.Default()

	// 允許來自所有來源的 CORS 請求
	// router.Use(cors.Default()) // 使用cors中介軟體

	// 設定路由
	router.Any("/go", func(c *gin.Context) {
		if c.Request.Method == "GET" {
			c.String(200, "Hello GO!")
		}
	})

	// 設定 app 路由
	app.SetupRouter(router)

	// 處理沒有定義的路徑
	router.NoRoute(func(c *gin.Context) {
		// 將請求轉發到埠 5002
		c.Redirect(302, "http://127.0.0.1:5002")
	})

	// 監聽埠
	router.Run("127.0.0.1:5001")
}
