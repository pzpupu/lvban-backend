package api

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"lvban/middleware"
)

func SetRouter(g *gin.Engine) {
	apiRouter := g.Group("/api")
	apiRouter.Use(gzip.Gzip(gzip.DefaultCompression)).Use(middleware.WxCloud())
	// 用户接口 依次: 分页列表，单条，插入，修改，删除
	//UserApi{}.init(apiRouter)
	// 自己的接口
	SelfApi{}.init(apiRouter)
	SettingApi{}.init(apiRouter)
}
