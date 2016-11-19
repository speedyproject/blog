package models

import (
	"blog/app/support"
	"fmt"
	"time"

	"github.com/russross/blackfriday"
)

import "encoding/json"

// Blogger model.
type Blogger struct {
	Id         int64     `xorm:"not null pk autoincr INT(11)"`
	Title      string    `xorm:"not null VARCHAR(50)"`
	Content    string    `xorm:"not null TEXT"`
	CategoryId string    `xorm:"'category_id' VARCHAR(20)"`
	Passwd     string    `xorm:"VARCHAR(64)"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreateBy   int       `xorm:"'create_by' not null INT(11)"`
	ReadCount  int64     `xorm:"'read_count' default 0 BIGINT(20)"`
	LeaveCount int64     `xorm:"'leave_count' default 0 BIGINT(20)"`
	Type       int
}

// Get blogger list.
func (b *Blogger) FindList() ([]Blogger, error) {
	// get list data from cache.
	list := make([]Blogger, 0)
	res, _ := support.Cache.Get(support.SPY_BLOGGER_LIST).Result()

	if res != "" {
		err := json.Unmarshal([]byte(res), &list)
		if err == nil {
			return list, err
		}
	}
	// if list data is null in cache,get list data in db.
	err := support.Xorm.Find(&list)

	if err == nil {
		res, e1 := json.Marshal(&list)
		if e1 != nil {
			support.Cache.Set(support.SPY_BLOGGER_LIST, string(res), 0)
		}
	}

	return list, err
}

//New to Add new blogger.
func (b *Blogger) New() (int64, error) {

	blog := new(Blogger)

	blog.Title = b.Title
	blog.Content = b.Content
	blog.CreateBy = b.CreateBy
	blog.CreateTime = time.Now()
	blog.Passwd = b.Passwd
	blog.CategoryId = b.CategoryId

	has, err := support.Xorm.InsertOne(blog)

	// refurbish cache.
	if err == nil {
		list := make([]Blogger, 0)
		err := support.Xorm.Find(&list)
		if err == nil {
			res, e1 := json.Marshal(&list)
			if e1 != nil {
				support.Cache.Set(support.SPY_BLOGGER_LIST, string(res), 0)
			}
		}
	}
	return has, err
}

// find blogger by id.
func (b *Blogger) FindById() (*Blogger, error) {

	blog := new(Blogger)
	// Get single blogger from cache.
	res, e1 := support.Cache.Get(support.SPY_BLOGGER_SINGLE + fmt.Sprintf("%d", b.Id)).Result()

	if e1 == nil {
		e2 := json.Unmarshal([]byte(res), &blog)
		if e2 == nil {
			return blog, nil
		}
	}
	// if cache not blogger data, find in db.
	_, err := support.Xorm.Id(b.Id).Get(blog)

	if err != nil {
		return blog, err
	}

	return blog, err
}

// Update blogger.
func (b *Blogger) Update() (bool, error) {

	has, err := support.Xorm.Id(b.Id).Update(&b)

	if err == nil {
		// refurbish cache.
		res, e1 := json.Marshal(&b)
		if e1 == nil {
			support.Cache.Del(support.SPY_BLOGGER_SINGLE + fmt.Sprintf("%d", b.Id))
			support.Cache.Set(support.SPY_BLOGGER_SINGLE+fmt.Sprintf("%d", b.Id), string(res), 0)
		}
	}

	return has > 0, err
}

// Delete blogger.
func (b *Blogger) Del() (bool, error) {

	has, err := support.Xorm.Id(b.Id).Delete(&b)

	if err == nil {
		// Delete cache.
		support.Cache.Del(support.SPY_BLOGGER_SINGLE + fmt.Sprintf("%d", b.Id))
	}

	return has > 0, err
}

func (b *Blogger) RenderContent() string {
	if b.Type == BLOG_TYPE_MD {
		return string(blackfriday.MarkdownCommon([]byte(b.Content)))
	}
	return b.Content
}
