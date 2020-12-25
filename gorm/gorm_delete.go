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
	db.LogMode(true)
	db.AutoMigrate(&User{}) // 同步迁移到数据库
	defer db.Close()

	// 查找对象进行删除
	var user User
	if db.First(&user, "name=?", "kk_3").Error == nil {
		db.Delete(&user)
	}

	// 批量删除
	db.Where("id>?", 3).Delete(&User{})

}
