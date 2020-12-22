package models

import "fmt"

const (
	dbUser     string = "root"
	dbPassword string = "Abcd1234_gome"
	dbHost     string = "10.112.76.35"
	dbPort     int    = 3306
	dbName     string = "todolist"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

//func main() {
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
//		dbUser, dbPassword, dbHost, dbPort, dbName)
//	DB, err := sql.Open("mysql", dsn)
//	defer DB.Close()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(-1)
//	}
//
//	if err := DB.Ping(); err != nil {
//		fmt.Println(err)
//		os.Exit(-1)
//	}
//
//}
