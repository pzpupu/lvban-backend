package model

// PlayTag 陪玩标签
type PlayTag struct {
	// 基础字段
	BaseModel
	// 标签名
	Name string `json:"name" form:"name" gorm:"uniqueIndex;size:64;not null"`
	// 所关联的陪玩
	//PlayCompanion []*PlayCompanion `gorm:"many2many:play_companion_tags;"`
}
