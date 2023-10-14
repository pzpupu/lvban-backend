package app

import (
	"lvban/common"
	"lvban/config"
	"os"
	"strconv"
)

func Run() {
	common.SetupGinLog()
	config.MysqlInit()
	var port = os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(*common.Port)
	}
	err := common.R.Run(":" + port)
	if err != nil {
		common.FatalLog("failed to start HTTP server: " + err.Error())
	}
}
