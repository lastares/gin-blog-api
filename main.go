package main

import (
	_ "go-gin-blog-api/docs"
	"go-gin-blog-api/initialize"
)

// @title 博客接口文档
// @version 1.0
// @description  博客接口文档
// @BasePath /api/v1/

func main() {
	//crontab.CronLaunch()
	// 初始化数据库连接
	initialize.MySQL()
	defer initialize.MysqlClose() // 延迟关闭数据库连接

	// 初始化Gin实例
	initialize.Gin()
}
