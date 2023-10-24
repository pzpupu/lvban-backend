package model

// PlayMedia 陪玩媒体表
type PlayMedia struct {
	// 基础字段
	BaseModel
	// 陪玩id
	PlayCompanionId uint `json:"play_companion_id" form:"play_companion_id" gorm:"index" `
	// 路径
	Path string `json:"path" form:"path" gorm:"uniqueIndex;size:255;not null"`
	// 媒体类型: 1=图片 2=视频 3=音频
	Type int `json:"type" form:"type" gorm:"not null"`
}
