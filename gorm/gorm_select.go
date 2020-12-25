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

	// 查询
	//for i := 0; i < 10; i++ {
	//	user := User{
	//		Name:        fmt.Sprintf("kk_%d", i),
	//		Password:    "123",
	//		Birthday:    time.Date(1988, 10, 12, 0, 0, 0, 0, time.UTC),
	//		Sex:         false,
	//		Tel:         "13000",
	//		Addr:        "四川",
	//		Description: "成都",
	//	}
	//	db.Create(&user)
	//}

	var user User

	db.First(&user, "name=?", "kk_6") // 拿到第一条数据
	fmt.Println(user)

	var user2 User // 拿到最后一条数据
	db.Last(&user2, "name=?", "kk_6")
	fmt.Println(user2)

	var users []User // 拿到所有数据
	db.Find(&users, "name=?", "kk_6")
	db.Where("name=?", "kk_6").Find(&users)
	fmt.Println(users)

	fmt.Println("======like=====")
	var like []User

	db.Where("name like ?", "kk_7").Find(&like)
	fmt.Println(like)

	fmt.Println("======in=====")
	var in []User

	db.Where("name in (?)", []string{"kk_6", "kk_8"}).Find(&in)
	fmt.Println(in)

	fmt.Println("查询多个值")

	var and []User
	db.Where("name=?", "kk_5").Where("password=?", "123").Find(&and)
	fmt.Println(and)

	fmt.Println("排序")

	var order []User
	db.Order("id desc").Find(&order)
	fmt.Println(order)

	fmt.Println("limit")
	var limit []User
	db.Limit(3).Offset(5).Order("id desc").Find(&limit)
	fmt.Println(limit)

	fmt.Println("查询总的数量")
	var count int
	db.Model(&User{}).Count(&count)
	fmt.Println(count)

	rows, _ := db.Model(&User{}).Select("name,password").Rows()
	for rows.Next() {
		var name, password string
		rows.Scan(&name, &password)
		fmt.Println(name, password)
	}

	rows, _ = db.Model(&User{}).Select("name,count(*)").Group("name").Rows()
	for rows.Next() {
		var name string
		var count int
		rows.Scan(&name, &count)
		fmt.Println(name, count)
	}

}
