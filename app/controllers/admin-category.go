package controllers

import "github.com/revel/revel"
import "blog/app/models"

// Category for blog, admin user can access
// 博客分类，后台用户操作
type Category struct {
	Admin
}

// Index page for manage category
// 后台管理博客分类的首页
func (c *Category) Index() revel.Result {
	return c.RenderTemplate("Admin/Category/Index.html")
}

// ListAll .
// 列出所有分类
func (c *Category) ListAll() revel.Result {
	category := new(models.Category)
	categorys := category.FindAll()
	return c.RenderJson(categorys)
}
