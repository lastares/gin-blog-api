package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"os"
)

func GetOffset(c *gin.Context) int {
	offset := 0
	page, _ := com.StrTo(c.Query("page")).Int()

	if page > 1 {
		offset = (page - 1) * com.StrTo(os.Getenv("PAGESIZE")).MustInt()
	}

	return offset
}
