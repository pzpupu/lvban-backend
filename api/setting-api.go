package api

import (
	"github.com/gin-gonic/gin"
	"lvban/common"
	"lvban/service"
	"net/http"
)

// SettingApi 控制器
type SettingApi struct {
}

func (t SettingApi) init(g *gin.RouterGroup) {
	group := g.Group("/setting")
	group.GET("/:key", t.get)
}

func (t SettingApi) get(g *gin.Context) {
	key := g.Param("key")
	g.JSON(http.StatusOK, common.OkData(service.SettingService.One(key).Value))
}
