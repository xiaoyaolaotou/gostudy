package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"todolist/utils"
)

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"`
	Tel      string    `json:"tel"`
	Addr     string    `json:"addr"`
	Desc     string    `json:"desc"`
}

// 验证密码
func (u *User) ValidatePassword(password string) bool {
	return utils.Md5(password) == u.Password
}

// 加载文件
func loadUsers() (map[int]User, error) {
	if file, err := ioutil.ReadFile("data/users.json"); err != nil {
		if os.IsNotExist(err) {
			return map[int]User{}, nil
		} else {
			return nil, err
		}
	} else {
		var users map[int]User
		if err := json.Unmarshal(file, &users); err == nil {
			return users, nil
		} else {
			return nil, err
		}
	}
}

func storeUsers(users map[int]User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("data/users.json", bytes, 0x066)
}

func GetUsers(q string) []User {
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	rtList := make([]User, 0)
	for _, user := range users {
		if q == "" || strings.Contains(user.Name, q) || strings.Contains(user.Addr, q) ||
			strings.Contains(user.Tel, q) || strings.Contains(user.Desc, q) {
			rtList = append(rtList, user)
		}
	}

	return rtList
}

func GetUserByName(name string) (User, error) {
	users, err := loadUsers()
	if err != nil {
		return User{}, err
	}
	for _, user := range users {
		if user.Name == name {
			fmt.Println(user.Name)
			return user, nil
		}
	}
	return User{}, fmt.Errorf("Not Foud")
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

func GetUserId() (int, error) {
	users, err := loadUsers()
	if err != nil {
		return -1, err
	}
	var id int
	for uid := range users {
		if id < uid {
			id = uid
		}
	}
	return id + 1, nil
}

func CreateUser(name, password, birthday, tel, addr, desc string) {
	id, err := GetUserId()
	if err != nil {
		panic(err)
	}
	day, _ := time.Parse("2006-01-02", birthday)
	user := User{
		Id:       id,
		Name:     name,
		Password: utils.Md5(password),
		Birthday: day,
		Tel:      tel,
		Addr:     addr,
		Desc:     desc,
	}
	users, err := loadUsers()
	if err != nil {
		panic(err)
	}
	users[id] = user
	storeUsers(users)
}
