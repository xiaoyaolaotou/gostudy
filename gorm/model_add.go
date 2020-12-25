package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

const (
	dbUser     string = "root"
	dbPassword string = "Abcd1234_gome"
	dbHost     string = "10.112.76.35"
	dbPort     int    = 3306
	dbName     string = "gorm"
)

type User struct {
	Id          int    `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(1024);unique"`
	Password    string
	Birthday    time.Time
	Sex         bool
	Tel         string
	Addr        string
	Description string
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	db.AutoMigrate(&User{}) // 同步迁移到数据库
	defer db.Close()

	// 添加单条数据
	user := User{
		Name:        "kk3",
		Password:    "123",
		Birthday:    time.Date(1988, 10, 12, 0, 0, 0, 0, time.UTC),
		Sex:         false,
		Tel:         "13000",
		Addr:        "四川",
		Description: "成都",
	}

	if db.NewRecord(user) { // 判断数据是否存在, 如果返回true， 则不存在， 可以创建
		db.Create(&user)

	}
}
