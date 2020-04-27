package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-gin-blog-api/docs"
	"go-gin-blog-api/orm"
	"go-gin-blog-api/routers"
)


// @title Swagger Example API
// @version 1.0
// @description gin swagger test.

// @host 127.0.0.1:8002
// @BasePath /v1/api/


func main() {
	// 数据库连接
	err := orm.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer orm.Close()

	// 实例化gin
	engine := gin.Default()

	config := &ginSwagger.Config{
		URL: "http://localhost:8002/swagger/doc.json",
	}
	//use ginSwagger middleware to
	engine.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	// 设置gin的运行模式
	gin.SetMode(gin.DebugMode)
	// 加载路由
	routers.InitRouter(engine)
	// 运行服务
	engine.Run(":8002")
}
