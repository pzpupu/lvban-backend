package model

// PlayCompanion 陪玩表
type PlayCompanion struct {
	// 基础字段
	BaseModel
	// 自增id
	//Id uint64 `json:"id" form:"id" gorm:"primaryKey;autoIncrement" `
	// 陪玩名称
	Name string `json:"name" form:"name" gorm:"size:64;not null"`
	// 性别 1:男 2:女
	Gender int8 `json:"gender" form:"gender" gorm:"type:tinyint(1);default:2;not null" `
	// 简介
	Summary string `json:"summary" form:"summary" gorm:"type:text" `
	// 陪玩媒体
	Medias []PlayMedia `json:"medias" gorm:"foreignKey:PlayCompanionId"`
	// 是否有语音介绍
	HasVoice bool `json:"hasVoice" form:"has_voice" gorm:"default:false" `
	// 陪玩项目
	Projects *[]PlayProject `json:"projects" gorm:"many2many:play_companion_projects;"`
	// 陪玩标签
	Tags *[]PlayTag `json:"tags" gorm:"many2many:play_companion_tag;foreignKey:Id;joinForeignKey:payId;joinReferences:tagId"`
	// 权重
	Weight int `json:"weight" form:"weight" gorm:"default:0" `
	// 时长
	//Duration []PlayDuration `gorm:"many2many:play_companion_durations;"`
	// 已售小时
	SoldHour int `json:"soldHour" form:"sold_hour" gorm:"default:0" `
	// 推荐
	IsRecommend bool `json:"isRecommend" form:"is_recommend" gorm:"default:false" `
}
