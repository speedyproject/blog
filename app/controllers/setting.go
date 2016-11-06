package controllers

import "github.com/revel/revel"
import "blog/app/models"

// Setting controller
type Setting struct {
	Admin
}

//SiteInfo model
type SiteInfo struct {
	Title      string
	SubTitle   string
	Url        string
	Seo        string
	Reg        string
	Foot       string
	Statistics string
	Status     string
}

// Site setting page.
func (s *Setting) SiteSetPage() revel.Result {

	set := new(models.Setting)
	data, err := set.FindAll()
	revel.INFO.Println(err)

	site := new(SiteInfo)

	if len(data) > 0 {
		for _, tmp := range data {
			switch tmp.Key {
			case "site-foot":
				site.Foot = tmp.Value
			case "site-reg":
				site.Reg = tmp.Value
			case "site-seo":
				site.Seo = tmp.Value
			case "site-status":
				site.Status = tmp.Value
			case "site-subtitle":
				site.SubTitle = tmp.Value
			case "site-title":
				site.Title = tmp.Value
			case "site-url":
				site.Url = tmp.Value
			case "site-statistics":
				site.Statistics = tmp.Value
			}
		}
	}
	s.RenderArgs["site"] = site
	return s.RenderTemplate("Admin/Setting/SiteSetPage.html")
}

//Site setting handler.
func (s *Setting) SiteSetHandler(title, subtitle, url, seo, reg, foot,
	statistics, status string) revel.Result {

	set := new(models.Setting)
	err := set.NewSiteInfo(title, subtitle, url, seo, reg, foot, statistics, status)

	if err != nil {
		return s.RenderJson("error")
	}

	return s.RenderJson("success")
}
