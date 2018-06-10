package controllers

type MessageController struct {
	BaseController
}

///message/nums/
// @router nums [post]
func (c *MessageController) Nums() {
	c.ToOKCount(0)
}
