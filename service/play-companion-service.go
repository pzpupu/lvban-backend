package service

import (
	"gorm.io/gorm"
	"lvban/model"
)

// setting 业务层
type playCompanionService struct {
	db *gorm.DB
}

// Page 分页查询
func (t playCompanionService) Page(page int, size int, v model.PlayCompanion, tag *int) (total int64, lists []model.PlayCompanion) {
	//playTag := model.PlayTag{BaseModel: model.BaseModel{Id: uint(*tag)}}

	// .Preload("Tags", "id = ", tag)
	condition := t.db.
		Debug().
		Model(&v).
		Joins("JOIN play_companion_tag ON play_companion.id = play_companion_tag.pay_id and play_companion_tag.tag_id = ?", tag).
		Where(&v)
	condition.Count(&total)
	if total <= 0 {
		return
	}

	condition.
		Preload("Medias").
		Preload("Projects").
		Preload("Tags").
		Offset((page - 1) * size).Limit(size).Find(&lists)
	return
}

// NewRecommend 新人推荐
func (t playCompanionService) NewRecommend(page int, size int) (total int64, lists []model.PlayCompanion) {
	//.Association("Languages")
	condition := t.db.Debug().Model(&model.PlayCompanion{}).Where("is_recommend = ?", true)
	condition.Count(&total)
	if total <= 0 {
		return
	}
	condition.
		Preload("Medias").
		Preload("Projects").
		Preload("Tags").
		Offset((page - 1) * size).Limit(size).Find(&lists)
	return
}
