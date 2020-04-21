package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-gin-blog-api/pkg/setting"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		dbType, dbName, user, password, host, tablePrefix string
	)

	database, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal("failed to get section 'database': %v", err)
	}

	dbType = database.Key("TYPE").String()
	dbName = database.Key("NAME").String()
	user = database.Key("USER").String()
	password = database.Key("PASSWORD").String()
	host = database.Key("HOST").String()
	tablePrefix = database.Key("TABLE_PREFIX").String()

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"),
		user,
		password,
		host,
		dbName,
	)

	if (err != nil) {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)

	db.LogMode(true)

	db.DB().SetMaxIdleConns(10)

	db.DB().SetMaxOpenConns(10)
}

func closeDb()  {
	defer db.Close()
}
