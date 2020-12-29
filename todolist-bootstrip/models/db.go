package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	dbUser     string = "root"
	dbPassword string = "Abcd1234_gome"
	dbHost     string = "10.112.76.35"
	dbPort     int    = 3306
	dbName     string = "todolist3"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", dsn)

	if err != nil || db.DB().Ping() != nil {
		panic("不能连接数据库")
	}
	db.AutoMigrate(&User{}, &Task{})
}
