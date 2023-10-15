package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func PrintHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header)

		// 请求前

		c.Next()

		// 请求后
		//latency := time.Since(t)
		//log.Print(latency)
		//
		//// 获取发送的 status
		//status := c.Writer.Status()
		//log.Println(status)
	}
}
