package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lvban/model"
	"lvban/utils"
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

// FindByOpenID 根据openid查询用户
func (t userService) FindByOpenID(id string) (*model.User, error) {
	user := model.User{}
	err := t.db.First(&user, "open_id = ?", id).Error
	return &user, err
}

// RegisterByOpenId 根据openid注册用户
func (t userService) RegisterByOpenId(openid string, data utils.UserInfoData) (*model.User, error) {
	// 检查openId是否已经存在
	var count int64
	t.db.Where("open_id = ?", openid).Count(&count)
	if count > 0 {
		return nil, errors.New("该用户已经注册")
	}
	// 创建用户
	user := model.User{OpenId: openid, Nickname: data.NickName, Gender: data.Gender, Avatar: data.AvatarUrl}
	result := t.db.Create(&user)
	return &user, result.Error
}
