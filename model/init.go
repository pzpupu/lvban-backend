package model

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
	"reflect"
	"time"
)

func InitDB() *gorm.DB {
	// 日志打印
	newLogger := logger.New(
		log.New(gin.DefaultWriter, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,         // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	source := "%s:%s@tcp(%s)/%s?readTimeout=15000ms&writeTimeout=15000ms&charset=utf8mb4&loc=Local&&parseTime=true"
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
	if dataBase == "" {
		dataBase = "lvban"
	}

	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	common.SysLog("start init mysql")
	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_", // 定义表前缀
			SingularTable: true, // true不在表后面+ s，
		},
	})

	common.SysLog("database connected")

	if err == nil {
		migrate := os.Getenv("DB_MIGRATE")
		if migrate != "" {
			common.SysLog("database begin migrate")
			for _, value := range []interface{}{
				&User{}, &PlayCompanion{}, &PlayMedia{}, &PlayProject{}, &PlayTag{}, &PlayDuration{}, &Setting{},
			} {
				err = db.AutoMigrate(value)
				if err != nil {
					t := reflect.TypeOf(value)
					if t.Kind() == reflect.Ptr {
						t = t.Elem()
					}
					common.FatalLog(fmt.Sprintf("failed to migrate table %s: %s", t.Name(), err.Error()))
				}
			}

			// 初始化系统默认配置
			InitSettingMap(db)
			common.SysLog("database migrated")

		}
	} else {
		common.FatalLog("DB Open error,err=", err.Error())
	}

	return db
}

type BaseModel struct {
	Id        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
