package models

import (
	"blog/app/support"
	"errors"

	"github.com/revel/revel"
)

var categoryModel *Category

// Category .
// 博客分类实体
type Category struct {
	Id     int64  `xorm:"not null pk autoincr INT(11)"`
	Name   string `xorm:"not null VARCHAR(20)"`
	Ident  string `xorm:"VARCHAR(30)"`
	Parent int64  `xorm:"default 0 INT(11)"`
	Desc   string `xorm:"VARCHAR(255)"`
}

// GetByIdent get category by category ident
// 通过 ident 获取分类
func (c *Category) GetByIdent(ident string) int64 {
	ca := &Category{}
	has, _ := support.Xorm.Where("ident = ?", ident).Get(ca)
	if has {
		return int64(ca.Id)
	}
	return 0
}

// 通过 ident 获取分类
func (c *Category) GetByID(id int64) (*Category, error) {
	ca := &Category{}
	has, _ := support.Xorm.Where("id = ?", id).Get(ca)
	if has {
		return ca, nil
	}
	return nil, errors.New("not found")
}

// AddOrUpdate is a function to save or update a category
// 添加或者更新一个分类
func (c *Category) AddOrUpdate(id int64, name, ident string, parent int64, desc string) (int64, error) {
	category := &Category{Name: name, Ident: ident, Parent: parent, Desc: desc}
	ca := &Category{}
	has, err := support.Xorm.Where("ident = ?", ident).Get(ca)
	if has && ca.Id != id && id != 0 {
		msg := "ident 已经存在"
		revel.ERROR.Println("save category errror: ", msg)
		return 0, errors.New(msg)
	}

	if id > 0 {
		support.Xorm.Id(id).Update(category)
		return id, nil
	}

	_, err = support.Xorm.Insert(category)
	if err != nil {
		revel.ERROR.Println("save category errror: ", err)
		return 0, err
	}
	return category.Id, nil
}

// Delete to delete a category
// 删除分类
func (c *Category) Delete(id int64) {
	c.resetOtherCategory(id)
	support.Xorm.Id(id).Delete(c)
}

// resetSubCategory to do:
// if a category is deleted, the parent of its child category
// would be set to 0
func (c *Category) resetOtherCategory(id int64) {
	sql := "UPDATE " + TABLE_CATEGORY + " SET `parent` = 0 WHERE id = ?"
	sql1 := "UPDATE " + TABLE_BLOG + "SET `category` = 0 WHERE category = ?"
	support.Xorm.Exec(sql, id)
	support.Xorm.Exec(sql1, id)
}

// RelatedBlogCount get how many blog that related to the category
// 获取该分类下的文章数目
func (c *Category) RelatedBlogCount() int {
	blogModel := new(Blog)
	count, err := support.Xorm.Where("category_id = ?", c.Id).Count(blogModel)
	if err != nil {
		revel.ERROR.Println("RelatedBlogCount error: ", err)
		return 0
	}
	return int(count)
}

// FindAll to find all categorys
// 查询所有的分类
func (c *Category) FindAll() *[]Category {
	categorys := make([]Category, 0)
	support.Xorm.Find(&categorys)
	return &categorys
}
