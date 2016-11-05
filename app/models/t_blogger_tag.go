package models

import (
	"blog/app/support"
)

//BloggerTag model
type BloggerTag struct {
	Type int    `xorm:"not null pk autoincr INT(11)"`
	Name string `xorm:"not null VARCHAR(20)"`
}

// Query all tag
func (b *BloggerTag) FindList() ([]BloggerTag, error) {
	bt := make([]BloggerTag, 0)
	err := support.Xorm.Find(&bt)
	return bt, err
}

// Add new tag
func (b *BloggerTag) New() (bool, error) {

	bt := new(BloggerTag)
	bt.Type = b.Type
	bt.Name = b.Name
	has, err := support.Xorm.InsertOne(&bt)

	return has > 0, err
}
