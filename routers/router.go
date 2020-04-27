package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/api/v1"
	"go-gin-blog-api/middleware/jwt"
	"net/http"
)

func InitRouter(engine *gin.Engine) *gin.Engine {
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	engine.GET("/auth", v1.GetAuth)
	apiV1 := engine.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tag/add", v1.AddTag)
		apiV1.POST("/tag/update", v1.UpdateTag)
		apiV1.DELETE("/tag/:id", v1.DeleteTag)

	}
	return engine
}