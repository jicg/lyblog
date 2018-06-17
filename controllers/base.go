package controllers

import (
	"github.com/astaxie/beego"

	"github.com/jicg/lyblog/models"
)

const USER_KEY = "USER"

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
	User    *models.User
	IsLogin bool
}

type Ret struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}
type RetH map[string]interface{}

func (ctx *BaseController) Prepare() {
	user := ctx.GetSession(USER_KEY)
	ctx.Data["IsLogin"] = false
	ctx.IsLogin = false
	if user != nil {
		ctx.User = user.(*models.User)
		ctx.Data["User"] = ctx.User
		ctx.IsLogin = true
		ctx.Data["IsLogin"] = true
	}
	ctx.Data["Title"] = "论坛"
	ctx.Data["Page"] = "index"
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
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
