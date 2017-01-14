package models

import (
	"blog/app/support"
	"fmt"
	"strings"

	"github.com/revel/revel"
)

//BloggerTag model
// 标签表
type Tag struct {
	Id     int64  `xorm:"not null pk autoincr INT(11)" json:"id"`
	Type   int    `xorm:"not null INT(11)"`
	Name   string `xorm:"not null VARCHAR(20)" json:"name"`
	Parent int64  `xorm:"default 0 INT(11)"`
	Ident  string `xorm:"VARCHAR(255)" json:"ident"`
}

// 标签关联表
type BlogTag struct {
	Id     int64 `xorm:"not null pk autoincr INT(11)"`
	Blogid int64 `xorm:"INT(11)"`
	Tagid  int64 `xorm:"INT(11)"`
}

// Query all tag
// 查找所有 tag
func (b *Tag) ListAll() ([]Tag, error) {
	bt := make([]Tag, 0)
	err := support.Xorm.Find(&bt)
	return bt, err
}

// 根据 id 获取标签
func (b *Tag) GetByID(id int64) (*Tag, error) {
	tag := new(Tag)
	has, err := support.Xorm.Id(id).Get(tag)
	if has {
		return tag, nil
	}
	return nil, err
}

// 根据 ident 获取标签
func (b *Tag) GetByIdent(ident string) (*Tag, error) {
	tag := &Tag{}
	has, err := support.Xorm.Where("ident = ?", ident).Get(tag)
	if has {
		return tag, nil
	}
	return nil, err
}

// New to Add a new tag
// 新增一个标签
func (b *Tag) New() (int64, error) {
	if b.Name != "" {
		bt := new(Tag)
		bt.Type = b.Type
		bt.Name = b.Name
		bt.Parent = b.Parent
		_, err := support.Xorm.InsertOne(bt)

		if err != nil {
			return 0, err
		}
		return bt.Id, nil
	}
	return 0, nil
}

// FindBlogCount to get count of blog related to this tag
// 查询标签关联的文章数目
func (t *Tag) FindBlogByTag(ident string) []Blog {
	var id int64
	if len(ident) > 0 {
		tag, err := t.GetByIdent(ident)
		if err != nil {
			id = tag.Id
		}
	} else {
		id = t.Id
	}
	sql := "SELECT b.* FROM " + TABLE_BLOG + " AS b, " + TABLE_TAG + " AS t, " + TABLE_BLOG_TAG + " AS bt WHERE b.id = bt.blogid AND t.id = bt.tagid AND t.id = " + fmt.Sprintf("%d", id)
	blogs := make([]Blog, 0)
	support.Xorm.Sql(sql).Find(&blogs)
	return blogs
}

// QueryTags to Search for tag
// 根据用户输入的单词匹配 tag
func (t *Tag) QueryTags(str string) ([]map[string][]byte, error) {
	sql := "SELECT name,id FROM t_tag WHERE name LIKE \"%" + str + "%\" ORDER BY LENGTH(name)-LENGTH(\"" + str + "\") ASC LIMIT 10"
	//sql := "SELECT name FROM t_tag"
	ress, err := support.Xorm.Query(sql)
	revel.TRACE.Println("res: ", ress)
	if err != nil {
		revel.ERROR.Println("err: ", err)
		return ress, err
	}
	return ress, nil
}

// 更新标签
func (t *Tag) Update() bool {
	if t.Id <= 0 {
		return false
	}
	support.Xorm.Id(t.Id).Update(t)
	return true
}

// 删除标签
func (t *Tag) Delete(ids []string) {
	revel.TRACE.Println("tags", ids)
	idStr := strings.Join(ids, ",")
	sql := "DELETE FROM " + TABLE_TAG + " WHERE id in (" + idStr + ")"
	sql2 := "DELET FROM " + TABLE_BLOG_TAG + " WHERE tagid in(" + idStr + ")"
	support.Xorm.Exec(sql)
	support.Xorm.Exec(sql2)
}

func (bt *BlogTag) AddTagRef(tagid, blogid int64) {
	btr := &BlogTag{Blogid: blogid, Tagid: tagid}
	support.Xorm.Insert(btr)
}
