package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	fmt.Println("homecontroler")
	c.Ctx.Output.Body([]byte("home"))
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.Output.Body([]byte("user: " + c.Ctx.Input.Param(":id")))
}

func main() {

	// 基本路由
	// 为根路径的get方法绑定函数
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hi, beego"))
	})

	beego.Post("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("post"))
	})

	beego.Any("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("any"))
	})

	// 提交参数
	// 正则表达式: 数据类型 \d
	// 任意多个 [0-9]*
	// 至少一个 [0-9]+
	beego.Any("/delete/:id(\\d{1,8})/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("delete: " + ctx.Input.Param(":id")))
	})

	beego.Any("/get/:id:int/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("get: " + ctx.Input.Param(":id")))
	})

	beego.Any("/put/:id([0-9a-zA-Z_]{4,16})/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("put: " + ctx.Input.Param(":id")))
	})

	beego.Router("/home/", &HomeController{})
	beego.Router("/user/:id(\\d+)/", &UserController{})

	beego.Run()
}
