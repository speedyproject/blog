package controllers

import "github.com/revel/revel"

type Search struct {
	*revel.Controller
}

func (s *Search) Index(q string) revel.Result {
	s.RenderArgs["keywords"] = q
	return s.RenderTemplate("Search/Index.html")
}
