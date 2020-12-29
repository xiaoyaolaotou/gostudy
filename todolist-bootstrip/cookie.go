package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:     "tag",
			Value:    "1",
			Path:     "/",
			MaxAge:   60,
			Expires:  time.Now().Add(time.Second * 60),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
	})

	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("tag")
		fmt.Println(cookie, err)
	})

	http.ListenAndServe(":8888", nil)
}
