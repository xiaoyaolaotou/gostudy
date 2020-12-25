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
	Name     string `gorm:""`
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

	// 添加索引
	// db.Model(&User{}).AddIndex("idx_name", "name")

	// 添加联合索引
	// db.Model(&User{}).AddIndex("idx_name_addr", "name", "addr")

	// 删除索引
	// db.Model(&User{}).RemoveIndex("idx_name_addr")

	// 添加唯一索引
	// db.Model(&User{}).AddUniqueIndex("idx_name", "name")

}
