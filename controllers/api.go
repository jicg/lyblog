package controllers

import (
	"github.com/jicg/lyblog/sysutils"
	"time"
	"fmt"
	"github.com/jicg/lyblog/models"
)

type ApiController struct {
	BaseController
}

// @router /activate [post]
func (c *ApiController) Activate() {
	if !c.IsLogin {
		c.Redirect("/user/login", 302)
		return
	}
	// 得到token
	token := sysutils.Token()
	email := c.User.Email
	//发送邮件
	if err := c.SendEmail("激活邮箱", email, token, "/user/activate"); err != nil {
		c.ToError(err.Error())
	}
	//将token保存在缓存里面，缓存30分钟
	c.cache.Put(email, token, 30*time.Minute)
	c.ToOK("邮件发送成功！")
}

// @router /template_email [get]
func (c *ApiController) ForgetEmail() {
	email := c.GetString("email", "")
	title := c.GetString("title", "找回密码")
	callurl := c.GetString("callurl", "/user/forget")
	if len(email) == 0 {
		c.CustomAbort(500, "请输入邮箱")
	}
	c.Data["email"] = email
	u, err := models.GetUserByEmail(email)
	if err != nil {
		c.CustomAbort(500, "输入的邮箱非法，不存在！")
	}
	token := c.GetString("token", "")
	c.Data["User"] = u
	c.Data["Title"] = title
	c.Data["Host"] = c.Ctx.Request.Host
	c.Data["AuthUrl"] = fmt.Sprintf("%s?email=%s&token=%s", c.RemoteUrl()+callurl, email, token)
	c.TplName = "common/template_email.html"
}
