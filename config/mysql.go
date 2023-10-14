package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"lvban/common"
	"os"
	"time"
)

func MysqlInit() {
	// 日志打印
	newLogger := logger.New(
		log.New(gin.DefaultWriter, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8mb4&loc=Local&&parseTime=true"
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
	if dataBase == "" {
		dataBase = "lvban"
	}

	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	common.SysLog("start init mysql")
	v, err := gorm.Open(mysql.Open(source), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 定义表前缀
			SingularTable: true, // true不在表后面+ s，
		},
	})

	if err != nil {
		if err != nil {
			common.FatalLog("DB Open error,err=", err.Error())
		}
	}
	common.DB = v

	common.SysLog("database connected")

}
