package controllers

import (
	"blog/app/models"

	"github.com/revel/revel"
)

var categoryModel models.Category

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
	c.ViewArgs["categorys"] = categorys
	return c.RenderTemplate("Admin/Category/Index.html")
}

// ListAll .
// 列出所有分类
func (c *Category) ListAll() revel.Result {
	category := new(models.Category)
	categorys := category.FindAll()
	return c.RenderJSON(categorys)
}

// EditPage is a page to edit a category
// 编辑分类页面
func (c *Category) EditPage(cid int64) revel.Result {
	ca, err := categoryModel.GetByID(cid)
	if err != nil {
		revel.ERROR.Printf("获取 id 为 %d 的分类失败。", cid)
		return c.NotFound("分类未找到")
	}
	c.ViewArgs["category"] = ca
	c.ViewArgs["allcategorys"] = ca.FindAll()
	return c.RenderTemplate("Admin/Category/Edit.html")
}

// AddPage page of add a category
// 添加一个分类的页面
func (c *Category) AddPage() revel.Result {
	c.ViewArgs["category"] = new(models.Category)
	c.ViewArgs["allcategorys"] = categoryModel.FindAll()
	return c.RenderTemplate("Admin/Category/Edit.html")
}

// Add to add a new category
// 添加一个新的分类
func (c *Category) Add(ca_name string, ca_ident string, ca_p, ca_id int, ca_desc string) revel.Result {
	c.Validation.Required(ca_name).Message("分类名必填")
	c.Validation.Required(ca_ident).Message("分类标识必填")
	if c.Validation.HasErrors() {
		return c.RenderJSON(&ResultJson{Success: false, Msg: c.Validation.Errors[0].Message, Data: ""})
	}
	flag, err := categoryModel.AddOrUpdate(int64(ca_id), ca_name, ca_ident, int64(ca_p), ca_desc)
	if flag == 0 {
		return c.RenderJSON(&ResultJson{Success: false, Msg: "添加分类失败: " + err.Error(), Data: ""})
	}
	return c.RenderJSON(&ResultJson{Success: true, Msg: "", Data: flag})
}

// Del a category
// 删除一个分类
func (c *Category) Del(id int) revel.Result {
	if id > 0 {
		categoryModel.Delete(int64(id))
		return c.RenderJSON(&ResultJson{Success: true})
	}
	return c.RenderJSON(&ResultJson{Success: false, Msg: "id 不存在"})
}
