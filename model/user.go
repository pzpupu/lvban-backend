package model

import (
	"database/sql/driver"
	"gorm.io/gorm"
)

// GenderType 定义一个类型表示性别。
type GenderType int8

// 定义 GenderType 的可能值。
const (
	Unknown GenderType = iota
	Male
	Female
)

// String 实现 fmt.Stringer 接口。
func (g GenderType) String() string {
	switch g {
	case Male:
		return "男"
	case Female:
		return "女"
	default:
		return "未知"
	}
}

// Scan 实现 sql.Scanner 接口，用于从数据库读取值。
func (g *GenderType) Scan(value interface{}) error {
	*g = GenderType(value.(int64)) // 将数据库中的整数值转换为 GenderType 类型
	return nil
}

// Value 实现 driver.Valuer 接口，用于写入数据库值。
func (g GenderType) Value() (driver.Value, error) {
	return int64(g), nil // 将 GenderType 值转换为数据库可存储的整数值
}

type User struct {
	// 基础字段
	gorm.Model
	// 微信Openid
	OpenId string `json:"-" form:"open_id" gorm:"uniqueIndex;size:64;not null"`
	// 昵称
	Nickname string `json:"nickname" form:"nickname" gorm:"not null"`
	// 头像
	Avatar string `json:"avatar" form:"avatar" gorm:"not null"`
	// 性别
	Gender GenderType `json:"gender" form:"gender" gorm:"not null"`
	// 手机
	Mobile string `json:"mobile" form:"mobile" `
	// 创建时间
	//CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"autoCreateTime" `
	// 更新时间
	//UpdatedAt time.Time `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime" `
}
