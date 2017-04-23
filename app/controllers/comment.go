package controllers

import (
	"blog/app/models"

	"github.com/revel/revel"
)

//Comment controller.
type Comment struct {
	*revel.Controller
}

//New comment.
func (c *Comment) NewComment(content, name string, blogid int64) revel.Result {
	res := &ResultJson{Success: false}
	if content == "" {
		res.Msg = "内容不能为空"
		return c.RenderJSON(res)
	}
	if blogid <= 0 {
		res.Msg = "博客不存在"
		return c.RenderJSON(res)
	}
	if name == "" {
		name = "default-name"
	}
	_comment := &models.Comment{BlogId: blogid, Name: name, Content: content}
	err := _comment.NewComment()
	if err != nil {
		res.Msg = err.Error()
		return c.RenderJSON(res)
	}
	res.Success = true
	return c.RenderJSON(res)
}

//Delete comment.
func (c *Comment) DelComment() revel.Result {
	return c.RenderText("ok")
}

//Modify comment.
func (c *Comment) ModifyComment() revel.Result {
	return c.RenderText("ok")
}
