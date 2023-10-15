package model

import (
	"database/sql/driver"
	"gorm.io/gorm"
)

// GenderType 定义一个类型表示性别。
type GenderType int8

// 定义 GenderType 的可能值。
const (
	Male GenderType = iota
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
	*g = GenderType(value.(int8)) // 将数据库中的整数值转换为 GenderType 类型
	return nil
}

// Value 实现 driver.Valuer 接口，用于写入数据库值。
func (g GenderType) Value() (driver.Value, error) {
	return int8(g), nil // 将 GenderType 值转换为数据库可存储的整数值
}

type User struct {
	// 基础字段
	gorm.Model
	// 主键
	//Id uint64 `json:"id" form:"id" gorm:"primaryKey;autoIncrement" `
	// 微信id
	WechatId string `json:"wechat_id" form:"wechat_id" gorm:"uniqueIndex;size:64"`
	// 昵称
	Nickname string `json:"nickname" form:"nickname" `
	// 性别
	Gender GenderType `json:"gender" form:"gender" `
	// 手机
	Mobile string `json:"mobile" form:"mobile" `
	// 创建时间
	//CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"autoCreateTime" `
	// 更新时间
	//UpdatedAt time.Time `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime" `
}
