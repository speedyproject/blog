package controllers

import (
	"github.com/revel/revel"
	"blog/app/models"
)

// Category for blog, admin user can access
// 博客分类，后台用户操作
type Category struct {
	Admin
}

// Index page for manage category
// 后台管理博客分类的首页
func (c *Category) Index() revel.Result {
	categoryStruct := new(models.Category)
	categorys := categoryStruct.FindAll()
	c.RenderArgs["categorys"] = categorys
	return c.RenderTemplate("Admin/Category/Index.html")
}

// ListAll .
// 列出所有分类
func (c *Category) ListAll() revel.Result {
	category := new(models.Category)
	categorys := category.FindAll()
	return c.RenderJson(categorys)
}

// AddPage page of add a category
// 添加一个分类的页面
func (c *Category) AddPage() revel.Result {
	ca := new(models.Category)
	c.RenderArgs["category"] = ca
	c.RenderArgs["allcategorys"] = ca.FindAll()
	return c.RenderTemplate("Admin/Category/Edit.html")
}

// Add to add a new category
// 添加一个新的分类
func (c *Category) Add(ca_name string, ca_ident string, p_ca int, ca_desc string) revel.Result {
	ca := new(models.Category)
	c.Validation.Required(ca_name).Message("分类名必填")
	c.Validation.Required(ca_ident).Message("分类标识必填")
	if c.Validation.HasErrors() {
		return c.RenderJson(&ResultJson{Success:false, Msg:c.Validation.Errors[0].Message, Data:""})
	}
	flag := ca.Add(ca_name, ca_ident, int64(p_ca), ca_desc)
	if flag == 0 {
		return c.RenderJson(&ResultJson{Success:false, Msg:"添加分类失败", Data:""})
	}
	return c.RenderJson(&ResultJson{Success:true, Msg:"", Data:flag})
}
