package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-blog-api/setting"
)

func GetPage(c *gin.Context) int {
	offset := 0

	page, _ := com.StrTo(c.Query("page")).Int()

	if page > 1 {
		offset = (page - 1) * setting.PageSize
	}

	return offset
}