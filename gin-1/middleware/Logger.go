package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d \n",
			param.ClientIP,
			param.TimeStamp,
			param.Method,
			param.Path,
			param.StatusCode,
		)
	})
}
