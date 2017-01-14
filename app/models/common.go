package models

import (
	"blog/app/support"
)

var TABLE_BLOG string
var TABLE_BLOG_TAG string
var TABLE_TAG string
var TABLE_CATEGORY string

func InitModel() {
	TABLE_BLOG = support.Xorm.TableInfo(new(Blog)).Name
	TABLE_TAG = support.Xorm.TableInfo(new(Tag)).Name
	TABLE_BLOG_TAG = support.Xorm.TableInfo(new(BlogTag)).Name
	TABLE_CATEGORY = support.Xorm.TableInfo(new(Category)).Name
}

func SyncDB() error {
	engine := support.Xorm
	InitModel()
	return engine.Sync2(new(Admin), new(AdminRole), new(Blog), new(Category), new(Comment), new(Setting), new(Tag), new(BlogTag))
}
