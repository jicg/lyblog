package controllers

import (
	"github.com/jicg/lyblog/models"
	"strings"
	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

// @router /login [get]
func (c *UserController) LoginPage() {
	verify, err := models.GetRandVerification()
	if err != nil {
		c.Abort("501")
		beego.Error(err)
	}
	c.SetSession("key_verification", verify.Value)
	c.Data["Code"] = verify.Code
	c.TplName = "user/login.html"
}

// @router /login [post]
func (c *UserController) Login() {
	vercode := c.GetString("vercode", "")
	if len(vercode) == 0 {
		c.ToError("验证码不能为空")
	}
	tvercode := ""
	if sessionvercode := c.GetSession("key_verification"); sessionvercode != nil {
		tvercode = sessionvercode.(string)
	}
	if !strings.EqualFold(vercode, tvercode) {
		c.ToError("验证码非法")
	}
	email := c.GetString("email", "guest")
	pass := c.GetString("pass", "guest")
	user, err := models.GetUser(email, pass)
	if err != nil {
		c.ToError("用户名或密码错误")
		beego.Error(err)
	}
	c.SetSession(USER_KEY, user)
	c.ToOK("登陆成功！", "/")
}

// @router /reg [get]
func (c *UserController) RegPage() {
	c.TplName = "user/reg.html"
}

// @router /reg [post]
func (c *UserController) Reg() {
	c.TplName = "user/login.html"
}
