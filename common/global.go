package common

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var R = gin.Default()

// DB 全局数据库操作实例
var DB *gorm.DB
