package controllers

import (
	"strconv"
	"github.com/jicg/lyblog/models"
)

type JieController struct {
	BaseController
}


func (this *JieController) NestPrepare() {
	this.Data["Page"] = "jie"
}

// @router /:id/ [get]
func (c *JieController) Get() {
	//page := c.Ctx.Input.Param(":page");
	idstr := c.Ctx.Input.Param(":id")
	var (
		id    int
		err   error
		note  *models.Note
		resps []*models.Replay
	)
	if id, err = strconv.Atoi(idstr); err != nil {
		c.Abort("404")
	}
	note, err = models.QueryNoteById(id)
	if err != nil {
		c.Abort("404")
	}
	if resps, err = models.QueryRespsByNoteIdAndPage(note.Id, 0, 10); err != nil {
		c.Abort("404")
	}
	note.Reps = resps
	c.Data["Note"] = note
	c.Data["Reps"] = resps
	c.TplName = "jie/detail.html"
}
