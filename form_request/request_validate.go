package form_request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gin-blog-api/global"
	"go-gin-blog-api/response"
)

func ValidFailed(c *gin.Context, err error)  {
	for _, err2 := range err.(validator.ValidationErrors) {
		response.FormValidateFailed(c, err2.Translate(global.Translator))
	}
}