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
		//UnionId := c.GetHeader("X-WX-UNIONID")
		AppId := c.GetHeader("X-WX-APPID")
		AccessToken := c.GetHeader("X-WX-CLOUDBASE-ACCESS-TOKEN")

		fmt.Printf("WxCloud -> Header: %s\n", c.Request.Header)

		c.Set(common.OpenId, OpenId)
		c.Set(common.AppId, AppId)
		c.Set(common.AccessToken, AccessToken)

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
