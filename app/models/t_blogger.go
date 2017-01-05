package models

import (
	"blog/app/support"
	"encoding/json"
	"fmt"
	"time"

	"log"

	"github.com/russross/blackfriday"
)

const (
	BLOG_STATUS_NORMAL  = 0 // 正常状态
	BLOG_STATUS_PENDING = 1 // 审核状态
	BLOG_TYPE_MD        = 0
	BLOG_TYPE_HTML      = 1
	PAGE_SIZE           = 10
	TABLE_BLOG          = "t_blogger"
	TABLE_BLOG_TAG      = "t_blog_tag"
)

// Blogger model.
// 博客实体
type Blogger struct {
	Id            int64     `xorm:"not null pk autoincr INT(11)"`
	Title         string    `xorm:"not null default '' VARCHAR(50)"`
	ContentHTML   string    `xorm:"not null TEXT 'content_html'"`
	CategoryId    int       `xorm:"INT(11) 'category_id'"`
	Passwd        string    `xorm:"VARCHAR(64)"`
	CreateTime    time.Time `xorm:"created"`
	CreateBy      int       `xorm:"not null INT(11)"`
	ReadCount     int64     `xorm:"default 0 BIGINT(20)"`
	LeaveCount    int64     `xorm:"default 0 BIGINT(20)"`
	UpdateTime    time.Time `xorm:"TIMESTAMP"`
	BackgroundPic string    `xorm:"VARCHAR(255)"`
	Type          int       `xorm:"INT(1)"`
	ContentMD     string    `xorm:"TEXT 'content_md'"`
	Summary       string    `xorm:"VARCHAR(255)"`
	Status        int       `xrom:"INT(11)"`
}

// FindList to Get blogger list.
// 获取所有博客
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

// 获取博客的标签
func (b *Blogger) BlogTags() []BloggerTag {
	sql := "SELECT t.* FROM " + TABLE_BLOG + " AS b, " + TABLE_TAG + " AS t, " + TABLE_BLOG_TAG + " AS bt WHERE b.id = bt.blogid AND t.id = bt.tagid AND b.id = " + fmt.Sprintf("%d", b.Id)
	tags := make([]BloggerTag, 0)
	support.Xorm.Sql(sql).Find(&tags)
	log.Println("err", tags)
	return tags
}

// GetBlogByPage .
// 根据页面获取博客
func (b *Blogger) GetBlogByPage(page int) ([]Blogger, error) {
	list := make([]Blogger, 0)
	start := (page - 1) * PAGE_SIZE
	err := support.Xorm.Limit(PAGE_SIZE, start).Find(&list)
	return list, err
}

// FindById to find blogger by id.
// 通过 id 查找博客
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

// GetSummary to cut out a part of blog content
// 获取一篇博客的摘要
// 如果没有摘要则截取文章开头 300 个字符
func (b *Blogger) GetSummary() string {
	if b.Summary != "" {
		return b.Summary
	}
	if len(b.ContentHTML) < 300 {
		return b.ContentHTML
	}
	return b.ContentHTML[0:300]
}

// MainURL return the url of the blog
// TODO:Laily it is can be set as id, category, ident and so on
func (b *Blogger) MainURL() string {
	return fmt.Sprintf("/article/%d", b.Id)
}

// FindByCategory .
// 查找某个分类下的博客
func (b *Blogger) FindByCategory(categoryID int64) (*[]Blogger, error) {
	blogs := make([]Blogger, 0)
	err := support.Xorm.Where("category_id = ?", categoryID).Find(&blogs)
	return &blogs, err
}

// IsMD to judge it is written by Markdown
// 判断这篇博客是否由 markdown 书写
func (b *Blogger) IsMD() bool {
	return b.Type == BLOG_TYPE_MD
}

// GetLatestBlog .
// 获取最新的博客
func (b *Blogger) GetLatestBlog(n int) []Blogger {
	blogs := make([]Blogger, 0)
	support.Xorm.Limit(n, 0).Find(&blogs)
	return blogs
}

// GetBlogCount .
// 获取博客总数
func (b *Blogger) GetBlogCount() int64 {
	blog := new(Blogger)
	total, err := support.Xorm.Where("is_deleted = 0 ").Count(blog)
	if err != nil {
		fmt.Println("get blog count error: ", err)
		return 0
	}
	return total
}

func (b *Blogger) RenderContent() string {
	if b.Type == BLOG_TYPE_MD && b.ContentHTML == "" {
		mdContent := string(blackfriday.MarkdownCommon([]byte(b.ContentMD)))
		return mdContent
	}
	return b.ContentHTML
}

// New to Add new blogger.
// 新建一个博客
func (b *Blogger) New() (int64, error) {
	blog := new(Blogger)
	blog.Title = b.Title
	blog.ContentHTML = b.ContentHTML
	blog.CreateBy = b.CreateBy
	blog.UpdateTime = time.Now()
	blog.Passwd = b.Passwd
	blog.CategoryId = b.CategoryId
	blog.Summary = b.Summary

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

// Update blogger.
// 更新博客
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

// Del to Delete blogger.
// 删除一篇博客
func (b *Blogger) Del() (bool, error) {

	has, err := support.Xorm.Id(b.Id).Delete(&b)

	if err == nil {
		// Delete cache.
		support.Cache.Del(support.SPY_BLOGGER_SINGLE + fmt.Sprintf("%d", b.Id))
	}

	return has > 0, err
}

// 更新浏览次数
func (b *Blogger) UpdateView(id int64) {
	blog := &Blogger{Id: id}
	blog, err := blog.FindById()
	if err == nil {
		support.Xorm.Table(blog).Id(id).Update(map[string]interface{}{"read_count": blog.ReadCount + 1})
	}
}
