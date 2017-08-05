package controllers

import (
	"blog/app/models"
	"blog/app/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/revel/revel"
)

var (
	promHttp   = promhttp.Handler()
	promResult = &PromHttpStruct{}
)

//Main controller.
type Main struct {
	*revel.Controller
}

// 构造一个 revel.Result
// 提供 golang 应用数据给 prometheus 的结构体
type PromHttpStruct struct{}

func (p *PromHttpStruct) Apply(req *revel.Request, resp *revel.Response) {
	promHttp.ServeHTTP(resp.Out, req.Request)
}

//SiteInfo model
type SiteInfo struct {
	Title     string
	SubTitle  string
	Copyright string
}

// Main page.
// 博客首页
func (m *Main) Main() revel.Result {
	set := new(models.Setting)
	set.Key = "site-title"
	title, _ := set.Get()
	set.Key = "site-subtitle"
	subtitle, _ := set.Get()
	set.Key = "site-foot"
	copyr, _ := set.Get()

	// 页面处理
	var p int
	m.Params.Bind(&p, "page")
	if p == 0 {
		p = 1
	}
	blogModel := new(models.Blog)
	blogs, err := blogModel.GetBlogByPage(p, 0)
	if err != nil {
		revel.ERROR.Println("加载博客列表失败: ", err)
		return m.RenderError(err)
	}
	hotBLogs := blogModel.GetHotBlog(5)

	pageStruct := new(service.BlogPager)

	m.ViewArgs["blogs"] = blogs
	m.ViewArgs["hotblogs"] = hotBLogs
	m.ViewArgs["thisPage"] = p
	m.ViewArgs["pager"] = pageStruct.GetPager(p)
	category := new(models.Category)
	m.ViewArgs["categorys"] = category.FindAll()
	info := &SiteInfo{Title: title, SubTitle: subtitle, Copyright: copyr}
	return m.Render(info)
}

// 某个分类下的博客
func (m *Main) Blog4Category(ca string) revel.Result {
	blog := new(models.Blog)
	category := new(models.Category)
	id := category.GetByIdent(ca)
	var blogs *[]models.Blog
	if id > 0 {
		blogs, _ = blog.FindByCategory(id)
	}
	m.ViewArgs["blogs"] = blogs
	return m.RenderTemplate("Main/Blog4Search.html")
}

func (m *Main) Debug() revel.Result {
	return promResult
}
