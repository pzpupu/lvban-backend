package main

import (
	"github.com/gin-gonic/gin"
	"lvban/api"
	"lvban/common"
	"lvban/service"
	"os"
	"strconv"
)

func main() {
	EnvId := os.Getenv("CBR_ENV_ID")
	common.SysLog("CBR_ENV_ID: " + EnvId)

	common.SetupGinLog()

	// 配置 Service
	service.Setup()

	// 关闭数据库
	defer service.CloseDB()

	var port = os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(*common.Port)
	}
	// Initialize HTTP server
	r := gin.Default()

	// Initialize Api routers
	api.SetRouter(r)

	err := r.Run(":" + port)
	if err != nil {
		common.FatalLog("failed to start HTTP server: " + err.Error())
	}
}
