package main

import (
	"github.com/astaxie/beego"
)

// 响应
/*
Ctx
 c.ctx
 c.ctx.writer
 c.ctx.output
*/

type ResponseController struct {
	beego.Controller
}

func (c *ResponseController) Test() {
	c.Data["user"] = "kk"
	c.TplName = "ResponseController/test.html"
}

func main() {
	beego.AutoRouter(&ResponseController{})
	beego.Run()
}
