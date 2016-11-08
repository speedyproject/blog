package controllers

import "github.com/revel/revel"

// User for User Controller
type Post struct {
	Admin
}

func (p *Post) Index() revel.Result {
	return p.RenderTemplate("Admin/Post/Index.html")
}
