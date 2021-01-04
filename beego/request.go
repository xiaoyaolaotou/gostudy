package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginForm struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

type RequestController struct {
	beego.Controller
}

func (c *RequestController) Header() {
	// 请求控制器和动作
	fmt.Println(c.GetControllerAndAction())
	// 请求头信息
	// 请求行
	fmt.Println(c.Ctx.Input.Method())
	c.Ctx.Output.Body([]byte("header"))

	var form LoginForm
	err := c.ParseForm(&form)
	fmt.Println(err, form)

	fmt.Println("===========")
	//fmt.Println(c.Ctx.Input.CopyBody(1024 * 1024))

	//file, hedader, err := c.GetFile("x")
	//fmt.Println(file, hedader, err)

	// 保存文件
	c.SaveToFile("x", "test.txt")

}

func main() {

	// 用户提交的数据
	/*
		Context => c.ctx
			请求数据 c.ctx
					c.ctx.request => http.request

		Controller
			Get Post...等

	*/

	beego.AutoRouter(&RequestController{})
	beego.Run()
}
