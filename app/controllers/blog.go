package controllers

import (
	"blog/app/models"
	"blog/app/routes"

	"github.com/revel/revel"
)

// Blogger controller
type Blog struct {
	*revel.Controller
}

// BloggerPage to display the blog detail.
// 显示博客详情
func (b Blog) BlogPage(ident string) revel.Result {
	blogModel := &models.Blog{Ident: ident}
	blog, err := blogModel.FindByIdent()
	if err != nil {
		revel.ERROR.Println("加载博客失败: ", err)
		return b.Redirect(routes.Main.Main())
	}
	b.RenderArgs["blog"] = blog
	settingModel := new(models.Setting)
	set, _ := settingModel.GetSiteInfo()
	revel.ERROR.Println("comment: ", set.Comment)
	b.RenderArgs["comment"] = set.Comment
	go blog.UpdateView(blog.Id)
	return b.Render()
}

// LatestBlogger get laster n blog
// 获取最新的 n 条博客
func (b *Blog) LatestBlogger() {
	n := 10
	blogModel := &models.Blog{}
	blogModel.GetLatestBlog(n)
}
