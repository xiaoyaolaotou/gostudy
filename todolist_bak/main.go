package main

import (
	"net/http"
	_ "todolist/controllers"
)

func main() {
	http.ListenAndServe(":9999", nil)
}
