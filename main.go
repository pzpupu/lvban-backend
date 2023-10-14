package main

import (
	"github.com/gin-gonic/gin"
	"lvban/api"
	"lvban/common"
	"lvban/model"
	"lvban/service"
	"os"
	"strconv"
)

func main() {
	common.SetupGinLog()
	// Initialize SQL Database
	db := model.InitDB()

	// 配置 Service
	service.Setup(db)

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
