package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

var (
	DB         *gorm.DB
	Gin        *gin.Engine
	Validate   *validator.Validate
	Translator ut.Translator
)
