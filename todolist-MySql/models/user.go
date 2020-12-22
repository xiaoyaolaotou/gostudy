package models

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"todolist/utils"
)

type User struct {
	Id         int
	Name       string
	Password   string
	Sex        bool
	Birthday   time.Time
	Tel        string
	Addr       string
	Desc       string
	CreateTime time.Time
}

// 验证密码
func (u *User) ValidatePassword(password string) bool {
	return utils.Md5(password) == u.Password
}

func GetUsers(q string) []User {
	//var users []User
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	var user User
	var newUser []User
	var sql string

	if q == "" { // 全部获取
		sql = "select id,name,birthday,sex,tel,addr,`desc`,create_time from todolist_user"
		rows, err := db.Query(sql)
		if err == nil {
			for rows.Next() {
				rows.Scan(&user.Id, &user.Name, &user.Password, &user.Birthday, &user.Sex, &user.Tel, &user.Addr, &user.CreateTime)
				newUser = append(newUser, user)
			}
			return newUser
		}
	} else {
		sql = fmt.Sprintf("select id, name, birthday,sex, tel, addr, `desc`, create_time from todolist_user "+
			"where name like '%%%s%%' or addr like '%%%s%%' or `desc` like '%%%s%%'", q, q, q)
		rows, err := db.Query(sql)
		if err == nil {
			for rows.Next() {
				rows.Scan(&user.Id, &user.Name, &user.Password, &user.Birthday, &user.Sex, &user.Tel, &user.Addr, &user.CreateTime)
				newUser = append(newUser, user)
			}
			return newUser
		}
	}
	return make([]User, 0)

}

func GetUserByName(name string) (User, error) {
	var user User
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return user, err
	}

	if err := db.Ping(); err != nil {
		return user, nil
	}
	defer db.Close()
	row := db.QueryRow("select name,password,birthday,sex,tel,addr,create_time from todolist_user where name=?", name)
	err = row.Scan(&user.Name, &user.Password, &user.Birthday, &user.Sex, &user.Tel, &user.Addr, &user.CreateTime)

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
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	sql := "select id,name,birthday,sex,tel,addr,`desc`,create_time from todolist_user where id=?"
	row := db.QueryRow(sql, id)

	var user User
	err = row.Scan(&user.Id, &user.Name, &user.Birthday, &user.Sex, &user.Tel, &user.Addr, &user.Desc, &user.CreateTime)
	if err != nil {
		return User{}, errors.New("not found")
	}
	return user, nil

}

func CreateUser(name, password, birthday, tel, addr, desc string) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	birParse, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		fmt.Println(err)
	}

	sql := "insert into todolist_user(name, password, birthday, tel, addr, `desc`,create_time) " +
		"values(?, md5(?), ?, ?, ?, ?,now())"
	_, err = db.Exec(sql, name, password, birParse, tel, addr, desc)
	if err != nil {
		fmt.Println("数据插入错误错误:", err)
	}
}
