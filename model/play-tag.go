package model

import "gorm.io/gorm"

// PlayTag 陪玩标签
type PlayTag struct {
	// 基础字段
	gorm.Model
	// 标签名
	Name string `json:"name" form:"name" `
}
