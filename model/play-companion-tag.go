package model

import (
	"gorm.io/gorm"
	"time"
)

type PlayCompanionTag struct {
	PlayId    int `gorm:"primaryKey"`
	TagId     int `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
