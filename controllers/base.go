package controllers

import (
	"github.com/astaxie/beego"

	"github.com/jicg/lyblog/models"
)

const USER_KEY = "USER"

type BaseController struct {
	beego.Controller
	user *models.User
}

type Ret struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Action string `json:"action"`
}

func (ctx *BaseController) Prepare() {
	user := ctx.GetSession(USER_KEY)
	if user != nil {
		ctx.user = user.(*models.User)
		ctx.Data["User"] = user
	}
	ctx.Data["Title"] = "论坛"
}

func (ctx *BaseController) ToError(msg string) {
	ctx.Data["json"] = &Ret{
		Status: 1, Msg: msg,
	}
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
