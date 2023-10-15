package api

import (
	"github.com/gin-gonic/gin"
	"lvban/common"
	"lvban/service"
	"net/http"
)

// SelfApi 控制器
type SelfApi struct {
}

func (t SelfApi) init(g *gin.RouterGroup) {
	group := g.Group("/self")
	group.GET("/", t.get)
}

// 分页列表
func (t SelfApi) get(c *gin.Context) {
	openId := c.GetString(common.OPEN_ID)
	user, err := service.UserService.FindByOpenID(openId)
	if err != nil {
		// 用户未注册
		c.JSON(http.StatusOK, common.Fail(common.UserNotRegister))
		return
	}
	c.JSON(http.StatusOK, common.OkData(user))
}
