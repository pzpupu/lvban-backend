package service

import (
	"gorm.io/gorm"
	"lvban/common"
	"lvban/model"
)

var db *gorm.DB

var UserService userService
var SettingService settingService
var PlayCompanionService playCompanionService

func Setup() {
	// Initialize SQL Database
	db = model.InitDB()
	UserService = userService{db}
	SettingService = settingService{db}
	PlayCompanionService = playCompanionService{db}
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		common.FatalLog("failed to get orm db: " + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		common.FatalLog("failed to close database: " + err.Error())
	}
}
