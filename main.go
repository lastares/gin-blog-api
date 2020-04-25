package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/orm"
	"go-gin-blog-api/routers"
)

func main() {
	// 数据库连接
	err := orm.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer orm.Close()

	// 实例化gin
	engine := gin.Default()
	// 设置gin的运行模式
	gin.SetMode(gin.DebugMode)
	// 加载路由
	routers.InitRouter(engine)
	// 运行服务
	engine.Run(":8002")
}
