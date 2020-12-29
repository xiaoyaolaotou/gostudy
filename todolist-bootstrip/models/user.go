package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	"todolist/utils"
)

type User struct {
	gorm.Model
	Name       string    `gorm:"type:varchar(32);not null;default:''"`
	Password   string    `gorm:"type:varchar(1024);not null;default:''"`
	Sex        bool      `gorm:"not null;default:false"`
	Birthday   time.Time `gorm:"type:date;not null"`
	Tel        string    `gorm:"type:varchar(32);not null;default:''"`
	Addr       string    `gorm:"type:varchar(512);not null;default:''"`
	Desc       string    `gorm:"column:description;type:varchar(1024);not null;default:''`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
}

// 验证密码
func (u *User) ValidatePassword(password string) bool {
	return utils.Md5(password) == u.Password
}

func GetUsers(q string) []User {

	var user []User
	if q == "" {
		db.Find(&user)
	} else {
		q = `%` + q + `%`
		db.Where("name like ?", q).Or("addr like ?", q).Find(&user)
	}
	return user

}

func GetUserByName(name string) (User, error) {
	var user User

	err := db.First(&user, "name=?", name).Error

	return user, err

}

func ValidateCreateUser(name, password, birthday, tel, addr, desc string) map[string]string {
	errors := map[string]string{}
	if len(name) > 12 || len(name) < 4 {
		errors["name"] = "名称长度必须在4~12之间"
	} else if _, err := GetUserByName(name); err == nil {
		errors["name"] = "名称重复"
	}
	return errors
}

func GetUserId(id int) (User, error) {

	var user User
	err := db.Find(&user, "id=?", id).Error
	return user, err

}

func CreateUser(name, password, birthday, tel, addr, desc string) {
	birParse, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		panic(err)
	}
	user := User{
		Name:       name,
		Password:   utils.Md5(password),
		Birthday:   birParse,
		Tel:        tel,
		Addr:       addr,
		Desc:       desc,
		CreateTime: time.Now(),
	}
	db.Create(&user)

}
