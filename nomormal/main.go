package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 提交数据常规则
	// 1. 一个名字对应一个值
	addr := "0.0.0.0:8888"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.FormValue("x"))     // formvalue ， 已经做了parse解析
		fmt.Println(r.PostFormValue("y")) // 接收body中的
		if file, fileheader, err := r.FormFile("file"); err == nil {
			fmt.Println(fileheader.Filename)
			io.Copy(os.Stdout, file)
		}
	})
	http.ListenAndServe(addr, nil)
}
