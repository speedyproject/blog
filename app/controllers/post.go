package controllers

import (
	"blog/app/models"
	"fmt"

	"github.com/revel/revel"
)

// User for User Controller
type Post struct {
	Admin
}

func (p *Post) Index() revel.Result {
	return p.RenderTemplate("Admin/Post/Index.html")
}

func (p *Post) NewPost(title, content, passwd string, user, labelID, tagID int64) revel.Result {
	fmt.Println(title)
	blog := &models.Blogger{Title: title, Content: content, Passwd: passwd}
	blog.New()
	return p.RenderJson(blog)
}
