package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/hello", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {
	c.TplName = "result.html"
}
