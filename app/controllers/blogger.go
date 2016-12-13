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

// BloggerPage to display the blog detail.
// 显示博客详情
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

// LatestBlogger get laster n blog
// 获取最新的 n 条博客
func (b *Blogger) LatestBlogger() {
	n := 10
	blogModel := &models.Blogger{}

}
