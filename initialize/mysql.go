package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"go-gin-blog-api/global"
	"os"
	"time"
)

func MySQL() {
	// 读取数据库相关配置
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
	)
	// 连接数据库
	var err error
	global.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		logrus.Error("Database connection failed, ", err)
		os.Exit(0)
	}

	// sql 调试模式
	global.DB.LogMode(true)

	// 设置table不是负数形式
	global.DB.SingularTable(true)
	// 表前缀
	dbPrefix := os.Getenv("DB_PREFIX")
	// 完成表名：表前缀+表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbPrefix + defaultTableName
	}
	// 数据库连接池
	global.DB.DB().SetMaxIdleConns(10)           // 设置空闲连接池中的最大连接数
	global.DB.DB().SetMaxOpenConns(100)          // 设置数据库连接最大打开数
	global.DB.DB().SetConnMaxLifetime(time.Hour) // 设置可重用连接的最长时间
	// 获取通用数据库接口
	global.DB.DB().Ping()
}

// 关闭数据库
func MysqlClose() {
	global.DB.Close()
}
