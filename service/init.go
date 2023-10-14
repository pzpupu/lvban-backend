package service

import "gorm.io/gorm"

var UserService userService

func Setup(db *gorm.DB) {
	UserService = userService{db}
}
