package controllers

import "github.com/revel/revel"

//Comment controller.
type Comment struct {
	*revel.Controller
}

//New comment.
func (c *Comment) NewComment() revel.Result {
	return c.RenderText("ok")
}

//Delete comment.
func (c *Comment) DelComment() revel.Result {
	return c.RenderText("ok")
}

//Modify comment.
func (c *Comment) ModifyComment() revel.Result {
	return c.RenderText("ok")
}
