package service

import (
	"github.com/joho/godotenv"
	"lvban/common"
)

func init() {
	// 加载 .env 文件
	if err := godotenv.Load("../.env"); err != nil {
		common.SysLog("Not find .env file, Use environment variables")
	}
	Setup()
}
