package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lvban/model"
)

// userService 业务层
type userService struct {
	db *gorm.DB
}

// List 分页列表
func (t userService) List(page, size int, v *model.User) gin.H {
	lists := make([]model.User, 0) // 结果
	t.db.Model(&model.User{}).Where(&v).Order("").Offset((page - 1) * size).Limit(size).Find(&lists)
	var total int64 // 统计
	t.db.Model(&v).Where(&v).Count(&total)
	h := gin.H{"list": lists, "total": total}
	return h
}

// One 根据主键查询
func (t userService) One(id interface{}) (v model.User) {
	t.db.Find(&v, id)
	return
}

// Update 修改记录
func (t userService) Update(v model.User) {
	t.db.Model(&v).Updates(v)
}

// Insert 插入记录
func (t userService) Insert(v model.User) {
	t.db.Save(&v)
}

// Delete 根据主键删除
func (t userService) Delete(ids []int) {
	t.db.Delete(model.User{}, ids)
}

func (t userService) FindByOpenID(id string) (model.User, error) {
	user := &model.User{}
	err := t.db.First(user, "open_id = ?", id).Error
	return *user, err
}
