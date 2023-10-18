package model

import (
	"fmt"
	"gorm.io/gorm"
	"lvban/common"
)

// RawString 直接返回原始的字符串内容，而不是加上双引号
type RawString string

func (rs RawString) MarshalJSON() ([]byte, error) {
	return []byte(rs), nil
}

// Setting 系统设置表
type Setting struct {
	Key   string    `json:"key" gorm:"primaryKey"`
	Value RawString `json:"value,string" gorm:"type:json"`
}

//

// InitSettingMap 初始化的Setting配置
func InitSettingMap(db *gorm.DB) {
	var OptionMap = make(map[string]RawString)
	// 首页轮播图
	OptionMap["HomeSwiper"] = RawString("[{\"img\": \"https://cdn-we-retail.ym.tencent.com/tsr/home/v2/banner1.png\",\"path\":\"/pages/promotion-detail/index?promotion_id=1\"},{\"img\": \"https://cdn-we-retail.ym.tencent.com/tsr/home/v2/banner2.png\",\"path\":\"/pages/promotion-detail/index?promotion_id=2\"},{\"img\": \"https://cdn-we-retail.ym.tencent.com/tsr/home/v2/banner3.png\",\"path\":\"/pages/promotion-detail/index?promotion_id=3\"},{\"img\": \"https://cdn-we-retail.ym.tencent.com/tsr/home/v2/banner4.png\",\"path\":\"/pages/promotion-detail/index?promotion_id=4\"},{\"img\": \"https://cdn-we-retail.ym.tencent.com/tsr/home/v2/banner5.png\",\"path\":\"/pages/promotion-detail/index?promotion_id=5\"}]")
	initDefaultsSettingToDatabase(OptionMap, db)
}

func initDefaultsSettingToDatabase(setting map[string]RawString, db *gorm.DB) {
	for s := range setting {
		tx := db.Where("`key` = ?", s).FirstOrCreate(&Setting{Key: s, Value: setting[s]})
		if tx.Error != nil {
			common.SysError(fmt.Sprintf("failed to create setting %s: %s", s, tx.Error.Error()))
		}
	}
}
