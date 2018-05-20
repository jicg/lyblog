package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	page := c.Ctx.Input.Param(":page");
	beego.Info(page)
	c.TplName = fmt.Sprintf("user/%s.html", page)
}
