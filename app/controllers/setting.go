package controllers

import (
	"github.com/revel/revel"
)

// Setting controller
type Setting struct {
	Admin
}

// Site setting page.
func (s *Setting) SiteSetPage() revel.Result {
	return s.RenderTemplate("Admin/Setting/SiteSetPage.html")
}

//SiteInfo model
type SiteInfo struct {
	Title      string
	SubTitle   string
	Url        string
	Seo        string
	Reg        int
	Foot       string
	Statistics string
	Status     int
}

func (s *Setting) SiteSetHandler() revel.Result {

	si := new(SiteInfo)
	s.Params.Bind(&si, "si")

	revel.INFO.Println(si)

	return s.RenderHtml("ok")
}
