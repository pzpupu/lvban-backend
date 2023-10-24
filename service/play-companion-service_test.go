package service

import (
	"lvban/model"
	"testing"
)

func TestInsert(t *testing.T) {
	var playMedias []model.PlayMedia
	playMedias = append(playMedias, model.PlayMedia{Path: "/play/1/1.jpg", Type: 1})
	playMedias = append(playMedias, model.PlayMedia{Path: "/play/1/2.mp4", Type: 2})
	playMedias = append(playMedias, model.PlayMedia{Path: "/play/1/3.mp3", Type: 3})

	var projects []model.PlayProject
	projects = append(projects, model.PlayProject{Name: "吃"})
	projects = append(projects, model.PlayProject{Name: "喝"})
	projects = append(projects, model.PlayProject{Name: "玩"})
	projects = append(projects, model.PlayProject{Name: "乐"})

	var tags []model.PlayTag
	tags = append(tags, model.PlayTag{Name: "外国小姐姐"})
	tags = append(tags, model.PlayTag{Name: "探店达人"})
	tags = append(tags, model.PlayTag{Name: "技能达人"})

	playCompanion := model.PlayCompanion{
		Name:        "陪玩1",
		Medias:      playMedias,
		HasVoice:    true,
		Projects:    &projects,
		Tags:        &tags,
		IsRecommend: true,
		Weight:      0,
		SoldHour:    0,
		Summary:     "简介",
		Gender:      2,
	}

	playCompanion2 := model.PlayCompanion{
		Name:     "陪玩2",
		Medias:   playMedias,
		HasVoice: true,
		//Projects:    &projects,
		//Tags:        &tags,
		IsRecommend: true,
		Weight:      0,
		SoldHour:    0,
		Summary:     "简介",
		Gender:      2,
	}
	var plays = []model.PlayCompanion{playCompanion, playCompanion2}
	tx := db.Debug().Create(&plays)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
}
