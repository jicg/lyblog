package controllers

import (
	"strconv"
	"github.com/jicg/lyblog/models"
)

type JieController struct {
	BaseController
}

// @router /:id/ [get]
func (c *JieController) Get() {
	//page := c.Ctx.Input.Param(":page");
	idstr := c.Ctx.Input.Param(":id")
	var (
		id   int
		err  error
		note *models.Note
	)
	if id, err = strconv.Atoi(idstr); err != nil {
		c.Abort("404")
	}
	note, err = models.QueryNoteById(id)
	if err != nil {
		c.Abort("404")
	}
	c.Data["Note"] = note
	c.TplName = "jie/detail.html"
}
