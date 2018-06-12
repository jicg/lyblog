package controllers

import (
	"github.com/jicg/lyblog/models"
	"strings"
	"github.com/astaxie/beego"
	"time"
)

type UserController struct {
	BaseController
}


func (c *UserController) Page() string {
	return "user"
}

// @router /set [get]
func (c *UserController) UserSet() {
	if !c.IsLogin {
		c.Redirect("/user/login", 302)
		return
	}
	c.TplName = "user/set.html"
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
	verify, err := models.GetRandVerification()
	if err != nil {
		c.Abort("501")
		beego.Error(err)
	}
	c.SetSession("key_verification", verify.Value)
	c.Data["Code"] = verify.Code
	c.TplName = "user/reg.html"
}

// @router /reg [post]
func (c *UserController) Reg() {
	//验证码
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

	//密码
	pass := c.GetString("pass", "")
	repass := c.GetString("repass", "")
	if len(pass) == 0 {
		c.ToError("密码不能为空！")
	}
	if pass != repass {
		c.ToError("两次密码输入不同！")
	}
	username := c.GetString("username", "")
	email := c.GetString("email")
	if len(email) == 0 {
		c.ToError("邮箱不能为空！")
	}
	if len(username) == 0 {
		c.ToError("昵称不能为空！")
	}
	user := &models.User{
		UserName: username,
		//邮箱
		Email: email,
		//密码
		Pass: pass,
		//头像地址
		Avatar: "/static/images/avatar/default.png",
		//用户角色 管理员1 社区之光2 该号已被封-1
		Auth: 0,
		//飞吻
		Experience: 0,
		//加入时间
		JoinTime: time.Now(),
	}
	if _, err := models.SaveUser(user); err != nil {
		c.ToError("用户注册失败！" + err.Error())
		beego.Error(err)
	}
	c.SetSession(USER_KEY, user)
	c.ToOK("注册成功！", "/")
}

// @router /logout [get]
func (c *UserController) Logout() {
	c.SetSession(USER_KEY, nil)
	c.Redirect("/user/login", 302)
}
