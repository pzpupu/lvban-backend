package model

import "gorm.io/gorm"

// PlayCompanion 陪玩表
type PlayCompanion struct {
	// 基础字段
	gorm.Model
	// 自增id
	//Id uint64 `json:"id" form:"id" gorm:"primaryKey;autoIncrement" `
	// 陪玩媒体
	Medias []PlayMedia
	// 是否有语音介绍
	HasVoice bool `json:"has_voice" form:"has_voice" gorm:"default:false" `
	// 陪玩项目
	Projects []PlayProject `gorm:"many2many:play_companion_projects;"`
	// 陪玩标签
	Tags []PlayTag `gorm:"many2many:play_companion_tags;"`
	// 权重
	Weight int `json:"weight" form:"weight" gorm:"default:0" `
	// 时长
	Duration []PlayDuration `gorm:"many2many:play_companion_durations;"`
	// 已售小时
	SoldHour int `json:"sold_hour" form:"sold_hour" gorm:"default:0" `
	// 简介
	Summary string `json:"summary" form:"summary" gorm:"type:text" `
	// 性别 1:男 2:女
	Gender int8 `json:"gender" form:"gender" gorm:"type:tinyint;size:1;default:2" `
}
