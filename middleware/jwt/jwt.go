package jwt

import (
	"github.com/gin-gonic/gin"
	http2 "go-gin-blog-api/response"
	"go-gin-blog-api/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		translateId := "ok"
		token := c.Query("token")
		if token == "" {
			translateId = "Access token failed."
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				translateId = "Token check failed."
			} else if time.Now().Unix() > claims.ExpiresAt { // Token expired.
				translateId = "Access token failed."

			}
		}

		if translateId != "ok" {
			c.JSON(http.StatusOK, http2.Failed{http2.ErrorCode, http2.Translate(translateId)})
			c.Abort()
			return
		}

		c.Next()
	}
}
