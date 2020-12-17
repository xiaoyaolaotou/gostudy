package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := "0.0.0.0:8888"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.Method)
		fmt.Fprintf(w, time.Now().Format("2006-01-02:15:04:05.000"))

		// 请求体
		fmt.Println("======post=====")
		io.Copy(os.Stdout, r.Body)
	})

	http.ListenAndServe(addr, nil)
}
