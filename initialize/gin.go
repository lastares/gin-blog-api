package initialize

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-gin-blog-api/global"
	"go-gin-blog-api/routers"
	"os"
)

func Gin() {
	// 初始化gin实例
	global.Gin = gin.Default()

	// 设置gin的运行模式
	gin.SetMode(gin.DebugMode)

	// 加载路由
	routers.InitRouter(global.Gin)

	Validate()

	// 运行服务
	// global.Gin.Run(":8002") // gin原始服务
	// 使用fvbock/endless来替换默认的ListenAndServe进行优雅重启
	serve := endless.NewServer(":8002", global.Gin)
	err := serve.ListenAndServe()

	if err != nil {
		logrus.Error("The servce start failed, ", err.Error())
		os.Exit(0)
	}
}
