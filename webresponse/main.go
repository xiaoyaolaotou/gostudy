package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hi")
	})

	http.ListenAndServe(addr, nil)
}
