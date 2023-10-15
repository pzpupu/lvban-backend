package model

import "gorm.io/gorm"

// PlayProject 陪玩项目
type PlayProject struct {
	// 基础字段
	gorm.Model
	// 项目名
	Name string `json:"name" form:"name" `
}
