package controllers

import "github.com/revel/revel"
import "blog/app/models"

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

	info := &SiteInfo{Title: title, SubTitle: subtitle, Copyright: copyr}
	return m.Render(info)
}
