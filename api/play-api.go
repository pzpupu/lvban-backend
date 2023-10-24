package api

import (
	"github.com/gin-gonic/gin"
	"lvban/common"
	"lvban/model"
	"lvban/service"
	"net/http"
)

// PlayApi 陪玩控制器
type PlayApi struct {
	Page int  `form:"page,default=1"`
	Size int  `form:"size,default=10"`
	Tag  *int `form:"tag"`
}

// init 初始化路由
func (p PlayApi) init(router *gin.RouterGroup) {
	group := router.Group("/play")
	group.GET("", p.page)
}

// page 分页列表
func (p PlayApi) page(c *gin.Context) {
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusOK, common.Fail(common.InvalidParameterValue))
	}
	// 新人推荐
	if p.Tag != nil && *p.Tag == -1 {
		total, lists := service.PlayCompanionService.NewRecommend(p.Page, p.Size)
		c.JSON(http.StatusOK, common.OkPage(total, lists))
		return
	}
	v := model.PlayCompanion{}
	// 添加标签条件
	if p.Tag != nil {
		v.Tags = &[]model.PlayTag{{BaseModel: model.BaseModel{Id: uint(*p.Tag)}}}
	}
	total, lists := service.PlayCompanionService.Page(p.Page, p.Size, v, p.Tag)
	c.JSON(http.StatusOK, common.OkPage(total, lists))
}
