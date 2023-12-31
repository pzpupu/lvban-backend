package api

import (
	"github.com/gin-gonic/gin"
	"lvban/common"
	"lvban/model"
	"lvban/service"
	"net/http"
)

func (t UserApi) init(g *gin.RouterGroup) {
	// 依次: 分页列表，单条，插入，修改，删除
	group := g.Group("/user")
	group.GET("/list", t.list) //不设置限制条件的画默认查询所有
	group.GET("/one", t.one)
	group.POST("/insert", t.insert)
	group.POST("/update", t.update)
	group.POST("/delete", t.delete)
}

// UserApi 控制器
type UserApi struct {
	Page int   `form:"page,default=1"`
	Size int   `form:"size,default=10"`
	Ids  []int `form:"ids"`
}

// 分页列表
func (t UserApi) list(c *gin.Context) {
	_ = c.Bind(&t)
	v := model.User{}
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, common.OkData(service.UserService.List(t.Page, t.Size, &v)))
}

// 根据主键Id查询记录
func (t UserApi) one(c *gin.Context) {
	var v model.User
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, common.OkData(service.UserService.One(v.Id)))
}

// 修改记录
func (t UserApi) update(c *gin.Context) {
	var v model.User
	_ = c.ShouldBindJSON(&v)
	service.UserService.Update(v)
	c.JSON(http.StatusOK, common.OkMsg("修改成功！"))
}

// 插入记录
func (t UserApi) insert(c *gin.Context) {
	var v model.User
	_ = c.ShouldBindJSON(&v)
	service.UserService.Insert(v)
	c.JSON(http.StatusOK, common.OkMsg("插入成功！"))
}

// 根据主键删除
func (t UserApi) delete(c *gin.Context) {
	_ = c.Bind(&t)
	service.UserService.Delete(t.Ids)
	c.JSON(http.StatusOK, common.OkMsg("删除成功！"))
}
