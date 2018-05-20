package controllers

import "github.com/astaxie/beego"

type CaseController struct {
	beego.Controller
}
func (c *CaseController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Year"] = "2017"
	c.TplName = "case/case.html"
}
