package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gin-blog-api/crontab"
	_ "go-gin-blog-api/docs"
	"go-gin-blog-api/orm"
	"go-gin-blog-api/routers"
)


// @title 博客接口文档
// @version 1.0
// @description  博客接口文档
// @BasePath /api/v1/

func main() {
	crontab.CronLaunch()

	// 数据库连接
	err := orm.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer orm.Close()

	// 实例化gin
	engine := gin.Default()

	//config := &ginSwagger.Config{
	//	URL: "http://localhost:8002/swagger/doc.json",
	//}
	//use ginSwagger middleware to
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置gin的运行模式
	gin.SetMode(gin.DebugMode)
	// 加载路由
	routers.InitRouter(engine)
	// 运行服务
	engine.Run(":8002")
}
