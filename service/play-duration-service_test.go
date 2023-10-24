package service

import (
	"lvban/model"
	"testing"
)

func TestInitData(t *testing.T) {
	var m = []model.PlayDuration{
		{Quantity: 3, Unit: 1},
		{Quantity: 1, Unit: 2},
	}
	tx := db.Create(&m)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	t.Log("插入成功")
}
