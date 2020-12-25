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
	gorm.Model
	Name     string
	Password string
	Birthday *time.Time
	Sex      bool
	Tel      string
	Addr     string
	Desc     string
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

	fmt.Println(db.HasTable(&User{})) // 判断表是否存在

	// db.Model(&User{}).ModifyColumn("score", "string") // 修改列
	// db.Model(&User{}).DropColumn("score") // 删除列

}
