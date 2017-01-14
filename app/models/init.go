package models

import (
	"blog/app/support"
)

var TABLE_BLOG string
var TABLE_BLOG_TAG string
var TABLE_TAG string
var TABLE_CATEGORY string

func init() {
	TABLE_BLOG = support.Xorm.TableInfo(new(Blogger)).Name
	TABLE_TAG = support.Xorm.TableInfo(new(BloggerTag)).Name
	TABLE_BLOG_TAG = support.Xorm.TableInfo(new(BloggerTagRef)).Name
	TABLE_CATEGORY = support.Xorm.TableInfo(new(Category)).Name
}
