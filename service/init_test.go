package service

import (
	"github.com/joho/godotenv"
	"lvban/common"
)

func init() {
	// 加载 .env 文件
	if err := godotenv.Load("../.env"); err != nil {
		common.FatalLog("Error loading .env file")
	}
	Setup()
}
