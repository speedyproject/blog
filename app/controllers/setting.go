package controllers

import (
	"blog/app/models"

	"github.com/revel/revel"
)

// Setting controller
type Setting struct {
	Admin
}

// Site setting page.
func (s *Setting) SiteSetPage() revel.Result {

	set := new(models.Setting)
	site, err := set.GetSiteInfo()
	revel.INFO.Println(err)

	s.RenderArgs["site"] = site
	return s.RenderTemplate("Admin/Setting/SiteSetPage.html")
}

//Site setting handler.
func (s *Setting) SiteSetHandler(title, subtitle, url, seo, reg, foot,
	statistics, status string) revel.Result {

	set := new(models.Setting)
	err := set.NewSiteInfo(title, subtitle, url, seo, reg, foot, statistics, status)

	if err != nil {
		return s.RenderJson(err.Error())
	}

	return s.RenderJson("success")
}
