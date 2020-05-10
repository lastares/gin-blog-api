package global

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	Gin *gin.Engine
)
