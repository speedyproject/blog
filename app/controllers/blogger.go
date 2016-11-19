package controllers

import (
	"blog/app/models"
	"log"

	"blog/app/routes"

	"github.com/revel/revel"
)

// Blogger controller
type Blogger struct {
	*revel.Controller
}

//Blogger page.
func (b Blogger) BloggerPage(id int64) revel.Result {
	blogModel := &models.Blogger{Id: id}
	blog, err := blogModel.FindById()
	if err != nil {
		log.Println("load blog error: ", err)
		return b.Redirect(routes.Main.Main())
	}
	b.RenderArgs["blog"] = blog
	return b.Render()
}
