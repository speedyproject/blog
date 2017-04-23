package controllers

import (
	"blog/app/models"
	"blog/app/support"

	"github.com/revel/revel"
)

type Search struct {
	*revel.Controller
}

func (s *Search) Index(q string) revel.Result {
	s.ViewArgs["keywords"] = q
	ids := support.FullTextSearch(q)
	blogs := make([]*models.Blog, 0)
	for _, v := range ids {
		blogModel := &models.Blog{Id: v}
		b, err := blogModel.FindById()
		if err == nil {
			blogs = append(blogs, b)
		}
	}
	s.ViewArgs["blogs"] = blogs
	s.ViewArgs["flag"] = "search"
	return s.RenderTemplate("Main/Blog4Search.html")
}
