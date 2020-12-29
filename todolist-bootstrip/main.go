package main

import (
	"net/http"
	_ "todolist/controllers"
)

func main() {

	http.ListenAndServe(":8888", nil)
}
