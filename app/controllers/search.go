package controllers

import (
	"blog/app/models"
	"blog/app/service"

	"github.com/revel/revel"
)

type Search struct {
	*revel.Controller
}

func (s *Search) Index(q string) revel.Result {
	s.RenderArgs["keywords"] = q
	ids := service.FullTextSearch(q)
	blogs := make([]*models.Blogger, 0)
	for _, v := range ids {
		blogModel := &models.Blogger{Id: v}
		b, err := blogModel.FindById()
		if err == nil {
			blogs = append(blogs, b)
		}
	}
	s.RenderArgs["blogs"] = blogs
	s.RenderArgs["flag"] = "search"
	return s.RenderTemplate("Main/Blog4Search.html")
}
