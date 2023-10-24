package model

// PlayProject 陪玩项目
type PlayProject struct {
	// 基础字段
	BaseModel
	// 项目名
	Name          string           `json:"name" form:"name" gorm:"uniqueIndex;size:64;not null"`
	PlayCompanion []*PlayCompanion `gorm:"many2many:play_companion_projects;"`
}
