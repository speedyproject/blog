package controllers

import "github.com/revel/revel"
import "blog/app/models"

type Category struct {
	Admin
}

func (c *Category) Index() revel.Result {
	category := new(models.Category)
	c.RenderArgs["categorys"] = category.FindAll()
	return c.RenderTemplate("Admin/Category/Index.html")
}
