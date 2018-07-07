package controllers

import (
	"github.com/jicg/lyblog/models"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	notes, err := models.QueryTopNotes()
	if err != nil {
		logs.Error(err)
		c.Abort("501")
	}
	c.Data["TopItems"] = notes
	c.Data["HotItems"] = notes
	c.TplName = "index.html"
}
