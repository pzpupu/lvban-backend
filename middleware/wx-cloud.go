package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lvban/common"
)

func WxCloud() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		OpenId := c.GetHeader("X-WX-OPENID")
		UnionId := c.GetHeader("X-WX-UNIONID")
		fmt.Printf("OpenID: %s, UnionID: %s\n", OpenId, UnionId)

		c.Set(common.OPEN_ID, OpenId)

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
