package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := "0.0.0.0:8888"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, time.Now().Format("2006-01-02:15:04:05.000"))
		// 1. 提交数据方式
		/*
			在URL中传递数据 url？argname=1&argage=2
		*/
		r.ParseForm()                // 解析参数
		fmt.Println(r.Form)          // 接收的参数类型都是string
		fmt.Println(r.Form.Get("x")) // 接收的参数类型都是string
		fmt.Println(r.Form["x"])     // 获取多个值

		// 2. 通过body提交数据
		/*
			body中的数据格式:
				1. application/json
				2. a=b&c=1
				3. multipart/form-data
		*/
		// Form 可以获取URL中参数也可以获取Body中参数

		// postfrom 只包含body中的数据， 不包含url中的参数
		fmt.Println(r.PostForm)

	})

	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := ioutil.ReadAll(r.Body)
		var j map[string]interface{}
		json.Unmarshal(ctx, &j)
		fmt.Println(j)

		fmt.Fprintln(w, time.Now().Format("2006-01-02:15:04:05.000"))
	})
	// 上传文件
	http.HandleFunc("/file/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024 * 1024) // 每次接收文件过程中最大使用内存
		// url?
		// body k=v
		// body 文件内容

		//x := r.MultipartForm.File["file"]
		//fmt.Println(x[0].Size)

		if fileHeaders, ok := r.MultipartForm.File["file"]; ok {
			for _, fileHeader := range fileHeaders {
				fmt.Println(fileHeader.Filename)
				fmt.Println(fileHeader.Size)
				nfile, _ := os.Create("./" + fileHeader.Filename)
				file, _ := fileHeader.Open()
				io.Copy(nfile, file)
				defer file.Close()
				nfile.Close()
			}

		}

	})

	http.ListenAndServe(addr, nil)
}
