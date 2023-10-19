package service

import (
	"gorm.io/gorm"
	"lvban/model"
)

// setting 业务层
type settingService struct {
	db *gorm.DB
}

// One 根据主键查询
func (t settingService) One(id string) (v model.Setting) {
	t.db.Find(&v, "`key` = ?", id)
	return
}

// ListByIds 列表查询
func (t settingService) ListByIds(keys []string) (settings []model.Setting) {
	t.db.Debug().Find(&settings, keys)
	return settings
}
