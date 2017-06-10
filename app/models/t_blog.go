package models

import (
	"blog/app/support"
	"encoding/json"
	"fmt"
	"time"

	"github.com/revel/revel"
	"github.com/russross/blackfriday"
)

const (
	BLOG_STATUS_NORMAL  = 0 // 正常状态
	BLOG_STATUS_PENDING = 1 // 审核状态
	BLOG_TYPE_MD        = 0
	BLOG_TYPE_HTML      = 1
	PAGE_SIZE           = 10
)

// Blogger model.
// 博客实体
type Blog struct {
	Id            int64     `xorm:"not null pk autoincr INT(11)"`
	Ident         string    `xorm:"not null VARCHAR(255)`
	Title         string    `xorm:"not null default '' VARCHAR(250)"`
	ContentHTML   string    `xorm:"not null TEXT 'content_html'"`
	CategoryId    int64     `xorm:"INT(11)"`
	Passwd        string    `xorm:"VARCHAR(64)"`
	CreateTime    time.Time `xorm:"TIMESTAMP"`
	CreateBy      int64     `xorm:"not null INT(11)"`
	ReadCount     int64     `xorm:"default 0 BIGINT(20)"`
	LeaveCount    int64     `xorm:"default 0 BIGINT(20)"`
	UpdateTime    time.Time `xorm:"TIMESTAMP" updated`
	BackgroundPic string    `xorm:"VARCHAR(255)"`
	Type          int       `xorm:"INT(1)"`
	Status        int       `xorm:"INT(11)"`
	ContentMD     string    `xorm:"TEXT 'content_md'"`
	Summary       string    `xorm:"VARCHAR(255)"`
	Pic           string    `xorm:"VARCHAR(200)"`
	IsDeleted     int       `xorm:"default 0 TINYINT(1)"`
}

// FindList to Get blogger list.
// 获取所有博客
func (b *Blog) FindList() ([]Blog, error) {
	// get list data from cache.
	list := make([]Blog, 0)
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
func (b *Blog) BlogTags() []Tag {
	sql := "SELECT t.* FROM " + TABLE_BLOG + " AS b, " + TABLE_TAG + " AS t, " + TABLE_BLOG_TAG + " AS bt WHERE b.id = bt.blogid AND t.id = bt.tagid AND b.id = " + fmt.Sprintf("%d", b.Id)
	tags := make([]Tag, 0)
	support.Xorm.Sql(sql).Find(&tags)
	return tags
}

// 获取博客的标签并转换成 json
func (b *Blog) BlogTagsJSON() string {
	tags := b.BlogTags()
	revel.ERROR.Println("tags ", tags)
	bytearr, err := json.Marshal(tags)
	if err != nil {
		revel.ERROR.Println(err)
		return ""
	}
	return string(bytearr)
}

// GetBlogByPage .
// 根据页面获取博客
func (b *Blog) GetBlogByPage(page int, pageSize int) ([]Blog, error) {
	if pageSize == 0 {
		pageSize = PAGE_SIZE
	}
	list := make([]Blog, 0)
	start := (page - 1) * pageSize
	err := support.Xorm.Desc("id").Limit(pageSize, start).Find(&list)
	return list, err
}

// GetBlogByPage .
// 根据页面获取博客
// FIXME:Laily 这里顺便查出作者，否则页面显示的时候再查询作者信息效率太差
func (b *Blog) GetBlogByPageAND(uid, category int64, page int, pageSize int) ([]Blog, error) {
	if pageSize == 0 {
		pageSize = PAGE_SIZE
	}
	temp := support.Xorm.Desc("id").Where("1=1")
	if uid > 0 {
		temp.And("create_by = ?", uid)
	}
	if category > 0 {
		temp.And("category_id = ?", category)
	}
	list := make([]Blog, 0)
	start := (page - 1) * pageSize
	err := temp.Limit(pageSize, start).Find(&list)
	return list, err
}

// FindById to find blogger by id.
// 通过 id 查找博客
func (b *Blog) FindById() (*Blog, error) {
	blog := new(Blog)
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

// FindByIdent to find blogger by ident.
// 通过 id 查找博客
func (b *Blog) FindByIdent() (*Blog, error) {
	blog := new(Blog)
	// Get single blogger from cache.
	res, e1 := support.Cache.Get(support.SPY_BLOGGER_SINGLE + fmt.Sprintf("%s", b.Ident)).Result()
	if e1 == nil {
		e2 := json.Unmarshal([]byte(res), &blog)
		if e2 == nil {
			return blog, nil
		}
	}
	// if cache not blogger data, find in db.
	_, err := support.Xorm.Where("ident=?", b.Ident).Get(blog)
	if err != nil {
		return blog, err
	}
	return blog, err
}

// GetSummary to cut out a part of blog content
// 获取一篇博客的摘要
// 如果没有摘要则截取文章开头 300 个字符
func (b *Blog) GetSummary() string {
	if b.Summary != "" {
		return b.Summary
	}
	if len(b.ContentHTML) < 300 {
		return b.ContentHTML
	}
	return b.ContentHTML[0:300]
}

// MainURL return the url of the blog
// 博客的链接
func (b *Blog) MainURL() string {
	return "/article/" + b.Ident
}

// Auther to get blog auther
// 获取作者
func (b *Blog) Auther() *Admin {
	userID := b.CreateBy
	user := new(Admin)
	user, _ = user.GetUserByID(int64(userID))
	return user
}

func (b *Blog) Category() *Category {
	categoryID := b.CategoryId
	category, err := categoryModel.GetByID(int64(categoryID))
	if err != nil {
		return &Category{Name: ""}
	}
	return category
}

// FindByCategory .
// 查找某个分类下的博客
func (b *Blog) FindByCategory(categoryID int64) (*[]Blog, error) {
	blogs := make([]Blog, 0)
	err := support.Xorm.Where("category_id = ?", categoryID).Find(&blogs)
	return &blogs, err
}

// IsMD to judge it is written by Markdown
// 判断这篇博客是否由 markdown 书写
func (b *Blog) IsMD() bool {
	return b.Type == BLOG_TYPE_MD
}

// GetLatestBlog .
// 获取最热门的博客
func (b *Blog) GetHotBlog(n int) []Blog {
	blogs := make([]Blog, 0)
	support.Xorm.Desc("read_count").Limit(n, 0).Find(&blogs)
	return blogs
}

// GetLatestBlog .
// 获取最新的博客
func (b *Blog) GetLatestBlog(n int) []Blog {
	blogs := make([]Blog, 0)
	support.Xorm.Desc("id").Limit(n, 0).Find(&blogs)
	return blogs
}

// GetBlogCount .
// 获取博客总数
func (b *Blog) GetBlogCount() int64 {
	blog := new(Blog)
	total, err := support.Xorm.Where("is_deleted = 0 ").Count(blog)
	if err != nil {
		revel.ERROR.Println("get blog count error: ", err)
		return 0
	}
	return total
}

func (b *Blog) RenderContent() string {
	if b.Type == BLOG_TYPE_MD && b.ContentHTML == "" {
		mdContent := string(blackfriday.MarkdownCommon([]byte(b.ContentMD)))
		return mdContent
	}
	return b.ContentHTML
}

// New to Add new blogger.
// 新建一个博客
func (b *Blog) New() (int64, error) {
	blog := new(Blog)
	blog.Title = b.Title
	if b.Ident == "" {
		b.Ident = b.Title
	}
	blog.Ident = b.Ident
	blog.ContentHTML = b.ContentHTML
	blog.ContentMD = b.ContentMD
	blog.CreateBy = b.CreateBy
	blog.UpdateTime = time.Now()
	blog.Passwd = b.Passwd
	blog.CategoryId = b.CategoryId
	blog.Summary = b.Summary
	blog.CreateTime = b.CreateTime

	_, err := support.Xorm.InsertOne(blog)

	// refurbish cache.
	if err == nil {
		list := make([]Blog, 0)
		err := support.Xorm.Find(&list)
		if err == nil {
			res, e1 := json.Marshal(&list)
			if e1 != nil {
				support.Cache.Set(support.SPY_BLOGGER_LIST, string(res), 0)
			}
		}
	}
	return blog.Id, err
}

// Update blogger.
// 更新博客
func (b *Blog) Update() (bool, error) {
	has, err := support.Xorm.Id(b.Id).Update(b)
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


// BatchDel to delete a batch of blog
// 删除一批博客
func (b *Blog) BatchDel(ids []int64) (bool, string) {
	idStr := ""
	for _, id := range ids {
		idStr += fmt.Sprintf(",%d", id)
	}
	sql := "DELETE FROM " + TABLE_BLOG + " WHERE id in (" + idStr[1:] + ")"
	res, err := support.Xorm.Exec(sql)
	if err != nil {
		revel.ERROR.Println("delete blog error: ", err)
		return false, err.Error()
	}
	rowsAffect, err1 := res.RowsAffected()
	if err1 != nil {
		revel.ERROR.Println("delete blog error1: ", err)
		return false, err.Error()

	}
	return rowsAffect >= 0, ""
}

// 更新浏览次数
func (b *Blog) UpdateView(id int64) {
	blog := &Blog{Id: id}
	blog, err := blog.FindById()
	if err == nil {
		support.Xorm.Table(blog).Id(id).Update(map[string]interface{}{"read_count": blog.ReadCount + 1})
	}
}

// 删除该博客关联的所有标签
func (b *Blog) DeleteAllBlogTags() error {
	bt := &BlogTag{Blogid: b.Id}
	_, err := support.Xorm.Delete(bt)
	if err != nil {
		return err
	}
	return nil
}
