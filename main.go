package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/pkg/setting"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HttpPort),
	//	Handler:        engine,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}

	//s.ListenAndServe()
	engine.Run(fmt.Sprintf(":%d", setting.HttpPort))
}
