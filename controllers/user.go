package controllers

import (
	"github.com/jicg/lyblog/models"
	"github.com/astaxie/beego"
	"time"
	"path"
	"os"
	"fmt"
	"github.com/jicg/lyblog/sysutils"
)

const (
	PATH_IMG_AVATAR = "data/asset/avatar/"
	URL_AVATAR      = "/asset/avatar/"
)

type UserController struct {
	BaseController
}

func (this *UserController) NestPrepare() {
	this.Data["Page"] = "user"
}

// @router /upload [post]
func (c *UserController) Upload() {
	if err := os.MkdirAll(PATH_IMG_AVATAR, 0777); err != nil {
		beego.Error(err)
		c.ToError(err.Error())
		return
	}

	if !c.IsLogin {
		c.ToError("请重写登陆")
		return
	}
	f, _, err := c.GetFile("file")
	defer f.Close()
	if err != nil {
		c.ToError(err.Error())
		return
	}
	filename := fmt.Sprintf("%d.avatar", 233*c.User.Id) //h.Filename
	fpath := path.Join(PATH_IMG_AVATAR, filename)

	err = c.SaveToFile("file", fpath)
	if err != nil {
		c.ToError(err.Error())
		return
	}
	c.ToOKH("修改成功！", RetH{
		"url": path.Join(URL_AVATAR, filename),
	})
}

// @router /set [get]
func (c *UserController) Set() {
	if !c.IsLogin {
		c.Redirect("/user/login", 302)
		return
	}
	c.TplName = "user/set.html"
}

// @router /set [post]
func (c *UserController) Update() {
	if !c.IsLogin {
		c.ToError("请重写登陆")
		return
	}
	if email := c.GetString("email", c.User.Email); len(email) != 0 {
		if c.User.Email != email {
			c.User.Activate = false
		}
		c.User.Email = email
	} else {
		c.ToError("邮箱不能为空")
		return
	}

	if name := c.GetString("username", c.User.UserName); len(name) != 0 {
		c.User.UserName = name
	} else {
		c.ToError("昵称不能为空")
		return
	}

	c.User.Sex, _ = c.GetInt("sex", c.User.Sex)

	if city := c.GetString("city", c.User.City); len(city) != 0 {
		c.User.City = city
	} else {
		c.ToError("城市不能为空")
		return
	}

	if sign := c.GetString("sign", c.User.Sign); len(sign) != 0 {
		c.User.Sign = sign
	} else {
		c.ToError("签名不能为空")
		return
	}

	if avatar := c.GetString("avatar", c.User.Avatar); len(avatar) != 0 {
		c.User.Avatar = avatar
	}

	//密码
	pass := c.GetString("pass", "")
	if len(pass) > 0 {
		repass := c.GetString("repass", "")
		if pass != repass {
			c.ToError("两次密码输入不同！")
		}
		c.User.Pass = pass
	}

	if _, err := models.UpdateUser(c.User); err != nil {
		c.ToError("修改失败：" + err.Error())
		return
	}

	c.ToOK("修改成功！")
}

// @router /login [get]
func (c *UserController) LoginPage() {
	c.InitVercode()
	c.TplName = "user/login.html"
}

// @router /login [post]
func (c *UserController) Login() {
	c.CheckVercode()
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
	c.InitVercode()
	c.TplName = "user/reg.html"
}

// @router /reg [post]
func (c *UserController) Reg() {
	c.CheckVercode()
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
	c.SetSession(USER_KEY, &models.User{})
	c.Redirect("/user/login", 302)
}

// @router /forget [get]
func (c *UserController) ForgetPage() {
	if c.IsLogin {
		c.Redirect("/user", 302)
	}
	c.InitVercode()
	var step = 1
	qemail := c.GetString("email", "")
	if len(qemail) > 0 {
		qtoken := c.GetString("token", "")
		u, err := models.GetUserByEmail(qemail);
		if err == nil && u != nil && u.Id != 0 && len(qtoken) > 0 {
			val := c.cache.Get(qemail)
			if val != nil {
				if token, ok := val.(string); ok && qtoken == token {
					step = 3
					c.Data["User"] = u
					c.SetSession(USER_KEY, u)
				} else {
					step = 2
				}
			}
		}
	}
	c.Data["step"] = step
	c.TplName = "user/forget.html"
}

// @router /forget [post]
func (c *UserController) Forget() {
	if c.IsLogin {
		c.Redirect("/user", 302)
	}
	//验证码
	c.CheckVercode()
	//判断邮箱是否存在
	email := c.GetString("email", "")
	if len(email) == 0 {
		c.ToError("请输入邮箱")
	}
	//向forget_email发请求获取，发送邮箱的html
	// 得到token
	token := sysutils.Token()
	if err := c.SendEmail("找回密码", email, token, "/user/forget"); err != nil {
		c.ToError(err.Error())
	}
	//将token保存在缓存里面，缓存30分钟
	c.cache.Put(email, token, 30*time.Minute)
	c.ToOK("邮件发送成功！", "/user/forget")
}

// @router /repass [post]
func (c *UserController) Repass() {
	c.CheckVercode()
	//密码
	pass := c.GetString("pass", "")
	if len(pass) == 0 {
		c.ToError("密码不能为空！")
	}
	repass := c.GetString("repass", "")
	if pass != repass {
		c.ToError("两次密码输入不同！")
	}
	c.User.Pass = pass
	if _, err := models.UpdateUser(c.User); err != nil {
		c.ToError("修改失败：" + err.Error())
	}
	c.ToOK("修改成功！", "/")
}

// @router /activate [get]
func (c *UserController) ActivatePage() {
	if !c.IsLogin {
		c.Redirect("/user/login", 302)
		return
	}

	qemail := c.GetString("email", "")
	if len(qemail) > 0 {
		qtoken := c.GetString("token", "")
		u, err := models.GetUserByEmail(qemail);
		if err == nil && u != nil && u.Id != 0 && len(qtoken) > 0 {
			val := c.cache.Get(qemail)
			if val != nil {
				if token, ok := val.(string); ok && qtoken == token {
					//step = 3
					//c.Data["User"] = u
					//c.SetSession(USER_KEY, u)
					c.User.Activate = true
					models.SaveUser(c.User)
					c.Data["User"] = c.User
				}
			}
		}
	}
	c.TplName = "user/activate.html"
}
