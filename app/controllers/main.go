package controllers

import (
	"blog/app/models"
	"log"

	"github.com/revel/revel"
)

//Main controller.
type Main struct {
	*revel.Controller
}

//SiteInfo model
type SiteInfo struct {
	Title     string
	SubTitle  string
	Copyright string
}

//Main page.
func (m *Main) Main() revel.Result {

	set := new(models.Setting)
	set.Key = "site-title"
	title, _ := set.Get()
	set.Key = "site-subtitle"
	subtitle, _ := set.Get()
	set.Key = "site-foot"
	copyr, _ := set.Get()

	blogModel := new(models.Blogger)
	blogs, err := blogModel.FindList()
	if err != nil {
		log.Println("load blog list error: ", err)
		return m.RenderError(err)
	}
	m.RenderArgs["blogs"] = blogs
	category := new(models.Category)
	m.RenderArgs["categorys"] = category.FindAll()
	info := &SiteInfo{Title: title, SubTitle: subtitle, Copyright: copyr}
	return m.Render(info)
}

func (m *Main) Blog4Category(ca string) revel.Result {
	blog := new(models.Blogger)
	category := new(models.Category)
	id := category.GetByIdent(ca)
	var blogs *[]models.Blogger
	if id > 0 {
		blogs, _ = blog.FindByCategory(id)
	}
	m.RenderArgs["blogs"] = blogs
	return m.RenderTemplate("Main/Blog4Category.html")
}
