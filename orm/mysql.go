package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	_ "github.com/joho/godotenv/autoload"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
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
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// 设置table不是负数形式
	DB.SingularTable(true)
	// 表前缀
	dbPrefix := os.Getenv("DB_PREFIX")
	// 完成表名：表前缀+表名
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return dbPrefix + defaultTableName;
	}
	return DB.DB().Ping()
}

// 关闭数据库
func Close() {
	DB.Close()
}
