package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseSuccessJson(c *gin.Context, data interface{}, total int)  {
	c.JSON(http.StatusOK, ResponseJson{Ok, data, Meta{total}})
}

func Success(c *gin.Context)  {
	c.JSON(http.StatusOK, Succeed{Ok})
}

func Translate(errorCode int) string {
	return ErrorMessageMap[errorCode]
}

func InvalidOperation(c *gin.Context, errorCode int)  {
	c.JSON(http.StatusOK, Failed{Fail, Translate(errorCode)})
}

func FormValidateFailed(c *gin.Context, errorMessage string)  {
	c.JSON(http.StatusOK, Failed{Fail, errorMessage})
	return
}
