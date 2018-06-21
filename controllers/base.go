package controllers

import (
	"github.com/astaxie/beego"

	"github.com/jicg/lyblog/models"
	"github.com/astaxie/beego/cache"
	"strings"
	"fmt"
	"time"
	"io/ioutil"
	"github.com/jicg/lyblog/sysutils"
	"github.com/astaxie/beego/httplib"
	"errors"
)

const (
	USER_KEY         = "USER"
	KEY_VERIFICATION = "key_verification"
)

var (
	bm cache.Cache
)

type NestPreparer interface {
	NestPrepare()
}
type BaseController struct {
	beego.Controller
	User    *models.User
	IsLogin bool
	cache   cache.Cache
}
type Ret struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Action string `json:"action,omitempty"`
	Count  int    `json:"count,omitempty"`
}
type RetH map[string]interface{}

func (ctx *BaseController) Prepare() {
	user := ctx.GetSession(USER_KEY)
	ctx.Data["IsLogin"] = false
	ctx.IsLogin = false
	ctx.cache = bm
	if user != nil {
		u := user.(*models.User)
		if u.Id != 0 {
			ctx.User = u
			ctx.Data["User"] = ctx.User
			ctx.IsLogin = true
			ctx.Data["IsLogin"] = true
		}
	}
	ctx.Data["Title"] = "论坛"
	ctx.Data["Page"] = "index"
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (ctx *BaseController) InitVercode() {
	verify, err := models.GetRandVerification()
	if err != nil {
		ctx.Abort("501")
		beego.Error(err)
	}
	ctx.SetSession(KEY_VERIFICATION, verify.Value)
	ctx.Data["Code"] = verify.Code
}
func (ctx *BaseController) CheckVercode() {
	//验证码
	vercode := ctx.GetString("vercode", "")
	if len(vercode) == 0 {
		ctx.ToError("验证码不能为空")
	}
	tvercode := ""
	if sessionvercode := ctx.GetSession(KEY_VERIFICATION); sessionvercode != nil {
		tvercode = sessionvercode.(string)
	}
	if !strings.EqualFold(vercode, tvercode) {
		ctx.ToError("验证码非法")
	}
}

func (ctx *BaseController) ToError(msg string) {
	ctx.Data["json"] = &Ret{
		Status: 1, Msg: msg,
	}
	beego.Error(msg)
	ctx.ServeJSON()
	ctx.StopRun()
}

func (ctx *BaseController) ToOK(msg string, actions ... interface{}) {
	action := ""
	if len(actions) >= 1 {
		if actions[0] != nil {
			if a, ok := actions[0].(string); ok {
				action = a
			}
		}
	}
	ctx.Data["json"] = &Ret{
		Status: 0, Msg: msg, Action: action,
	}
	ctx.ServeJSON()
	ctx.StopRun()
}

func (ctx *BaseController) ToOKH(msg string, rets RetH) {
	rets["status"] = 0;
	rets["msg"] = msg;
	ctx.Data["json"] = rets
	ctx.ServeJSON()
	ctx.StopRun()
}

func (ctx *BaseController) ToOKCount(count int) {

	ctx.Data["json"] = &Ret{
		Status: 0, Count: count,
	}
	ctx.ServeJSON()
	ctx.StopRun()
}

func (ctx *BaseController) RemoteUrl() string {
	//token := sysutils.Token()
	scheme := "http://"
	if ctx.Ctx.Request.TLS != nil {
		scheme = "https://"
	}
	///user/forget_email?email=%s&token=%s
	url := fmt.Sprintf("%s%s", scheme, ctx.Ctx.Request.Host)
	return url
}

func (c *BaseController) SendEmail(title, email, token, callurl string) error {
	beego.Info("0---  ", time.Now().Unix())
	//向forget_email发请求获取，发送邮箱的html
	url := fmt.Sprintf("%s/api/template_email?title=%s&email=%s&token=%s&callurl=%s", c.RemoteUrl(), title, email, token, callurl)
	resp, err := httplib.Get(url).DoRequest()
	if err != nil {
		return err
	}
	status := resp.StatusCode
	bs, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if status != 200 {
		return errors.New(string(bs))
	}
	beego.Info("1---  ", time.Now().Unix())
	//发送邮件
	err = sysutils.SendEmail(sysutils.NewEmail(email, "找回密码", string(bs), "html"))
	beego.Info("2---  ", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func init() {
	//https://beego.me/docs/module/cache.md
	var err error
	bm, err = cache.NewCache("file", `{"CachePath":"data/.cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	if err != nil {
		beego.Error(err)
		return
	}
}
