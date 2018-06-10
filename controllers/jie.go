package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type JieIndexController struct {
	beego.Controller
}

func (c *JieIndexController) Get() {
	page := c.Ctx.Input.Param(":page");
	c.Data["Year"] = "2017"
	c.TplName = fmt.Sprintf("jie/%s.html", page)
}
