package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/api/v1"
	"net/http"
)

func InitRouter(engine *gin.Engine) *gin.Engine {
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	apiV1 := engine.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tag/add", v1.AddTag)
		apiV1.POST("/tag/update", v1.UpdateTag)
		apiV1.DELETE("/tag/:id", v1.DeleteTag)
	}
	return engine
}