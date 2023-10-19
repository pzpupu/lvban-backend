package api

import (
	"github.com/gin-gonic/gin"
	"lvban/common"
	"lvban/service"
	"net/http"
	"strings"
)

// SettingApi 控制器
type SettingApi struct {
}

func (t SettingApi) init(g *gin.RouterGroup) {
	group := g.Group("/setting")
	group.GET("/:key", t.get)
}

// get 获取配置
func (t SettingApi) get(g *gin.Context) {
	key := g.Param("key")
	ids := strings.Split(key, ";")
	settings := service.SettingService.ListByIds(ids)
	// 将 list 转换成map
	var jsonMap = make(map[string]interface{})
	for _, setting := range settings {
		jsonMap[setting.Key] = setting.Value
	}
	g.JSON(http.StatusOK, common.OkData(jsonMap))
}
