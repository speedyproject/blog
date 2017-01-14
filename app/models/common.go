package models

import (
	"blog/app/support"
)

func SyncDB() error {
	engine := support.Xorm
	return engine.Sync2(new(Admin), new(AdminRole), new(Blogger), new(Category), new(Comment), new(Setting), new(BloggerTag), new(BloggerTagRef))
}
