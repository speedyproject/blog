package models

import (
	"blog/app/support"

	"fmt"

	"github.com/revel/revel"
)

type Category struct {
	Id     int64  `xorm:"not null pk autoincr INT(11)`
	Name   string `xorm:"not null VARCHAR(15)"`
	Ident  string `xorm:"not null VARCHAR(15)"`
	Parent int64  `xorm:"not null INT(11)"`
}

// Add function to save a category
func (c *Category) Add(name, ident string, parent int64) int64 {
	category := &Category{Name: name, Ident: ident, Parent: parent}
	_, err := support.Xorm.Insert(category)
	if err != nil {
		revel.ERROR.Println("save category errror: ", err)
		return 0
	}
	return category.Id
}

// Delete to delete a category
func (c *Category) Delete(id int64) {
	resetSubCategory(id)
	support.Xorm.Id(id).Delete(c)
}

// resetSubCategory to do:
// if a category is deleted, the parent of its child category
// would be set to 0
func resetSubCategory(id int64) {
	sql := "UPDATE " + TABLE_CATEGORY + " SET `parent` = 0 WHERE id = ?"
	support.Xorm.Exec(sql, id)
}

//FindAll to find all categorys
func (c *Category) FindAll() *[]Category {
	categorys := make([]Category, 0)
	support.Xorm.Find(&categorys)
	fmt.Println("categorys: ", categorys)
	return &categorys
}
