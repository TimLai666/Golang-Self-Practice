// app.go

package app

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	// 設定 app 相關路由
	router.GET("/app", func(c *gin.Context) {
		c.String(200, "Hello app!")
	})
}
