package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lvban/common"
	"lvban/service"
	"lvban/utils"
	"lvban/vo"
	"net/http"
)

// SelfApi 控制器
type SelfApi struct {
}

func (t SelfApi) init(g *gin.RouterGroup) {
	group := g.Group("/self")
	group.GET("", t.get)
	group.POST("/register", t.post)
}

// get 获取当前用户信息
func (t SelfApi) get(c *gin.Context) {
	openId := c.GetString(common.OpenId)
	user, err := service.UserService.FindByOpenID(openId)
	if err != nil {
		// 用户未注册
		c.JSON(http.StatusOK, common.Fail(common.UserNotRegister))
		return
	}
	c.JSON(http.StatusOK, common.OkData(&user))
}

// post 为当前用户注册
func (t SelfApi) post(c *gin.Context) {
	//appid := c.GetString(common.AppId)
	openId := c.GetString(common.OpenId)
	//accessToken := c.GetString(common.AccessToken)
	json := &vo.CloudId{}
	if err := c.ShouldBindJSON(json); err != nil {
		// 请求的 JSON 格式错误，返回一个错误响应
		c.JSON(http.StatusOK, common.FailMsg(common.RequiredParameterIsEmpty, "CloudId为空"))
		return
	}

	data, err := utils.GetOpenData(openId, json.CloudId)
	if err != nil {
		common.SysError(fmt.Sprintf("获取 cloudID 对应的数据 -> openid: %s error: %s", openId, err.Error()))
		c.JSON(http.StatusOK, common.FailMsg(common.OpenDataFail, err.Error()))
		return
	}

	// 注册用户
	user, err := service.UserService.RegisterByOpenId(openId, data.Data)
	if err != nil {
		common.SysError(fmt.Sprintf("注册用户 -> openid: %s error: %s", openId, err.Error()))
		c.JSON(http.StatusOK, common.FailMsg(common.UserError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.OkData(user))
}
