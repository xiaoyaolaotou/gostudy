package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	dbUser     string = "root"
	dbPassword string = "Abcd1234_gome"
	dbHost     string = "10.112.76.35"
	dbPort     int    = 3306
	dbName     string = "todolist"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// 操作MySQL
	rows, err := db.Query("select name,password from todolist_user")
	var (
		name     string
		password string
	)

	for rows.Next() {
		rows.Scan(&name, &password)
		fmt.Println(name, password)
	}

}
