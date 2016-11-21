package controllers

import "github.com/revel/revel"
import "blog/app/service"
import "blog/app/models"
import "fmt"

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
	fmt.Println("blogs are: ", blogs)
	s.RenderArgs["blogs"] = blogs
	return s.RenderTemplate("Search/Index.html")
}
