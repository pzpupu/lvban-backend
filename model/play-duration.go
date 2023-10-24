package model

// PlayDuration 陪玩时长
type PlayDuration struct {
	// 基础字段
	BaseModel
	// 数量
	Quantity int8 `json:"quantity" form:"quantity" gorm:"default:1;index:,unique,composite:unindex" `
	// 单位: 1-小时, 2-天
	Unit int8 `json:"unit" form:"unit" gorm:"default:1;index:,unique,composite:unindex" `
}
