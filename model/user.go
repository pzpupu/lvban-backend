package model

type User struct {
	// 基础字段
	BaseModel
	// 微信Openid
	OpenId string `json:"-" form:"open_id" gorm:"uniqueIndex;size:64;not null"`
	// 昵称
	Nickname string `json:"nickname" form:"nickname" gorm:"not null"`
	// 头像
	Avatar string `json:"avatar" form:"avatar" gorm:"not null"`
	// 性别
	Gender int8 `json:"gender" form:"gender" gorm:"not null"`
	// 手机
	Mobile string `json:"mobile" form:"mobile" `
	// 创建时间
	//CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"autoCreateTime" `
	// 更新时间
	//UpdatedAt time.Time `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime" `
	// TODO: 用户来源(酒吧、线下场所等..)
}
