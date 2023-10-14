package api

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetRouter(g *gin.Engine) {
	apiRouter := g.Group("/api")
	apiRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	// 用户接口 依次: 分页列表，单条，插入，修改，删除
	UserApi{}.init(apiRouter)
}
