package controllers

import (
	"blog/app/models"
	"strconv"
	"strings"
	"time"

	"github.com/revel/revel"
)

/**
 * Add a blog for admin user
 * 发布博客 action
 */

var blogModel *models.Blog

// PostData model.
// 发布博客前端提交的数据
type PostData struct {
	Id          int64
	Title       string // 博客标题
	Ident       string // 博客标示，用作url
	ContentMD   string // 博客内容 MD
	ContentHTML string // 博客内容 HTML
	Category    int64  // 博客类别
	Tag         string // 标签 格式：12,14,32
	Keywords    string // 关键词 格式：java,web开发
	passwd      string // 博客内容是否加密
	Summary     string // 博客摘要
	Type        int    // 0 表示 markdown，1 表示 html
	NewTag      string // 新添加的标签
	Createtime  string // 创建时间
}

// User for User Controller
type Post struct {
	Admin
}

// Index page to create or edit a blog
// 创建博客页面，编辑页面也是这个
func (p *Post) Index(postid int64) revel.Result {
	categoryModel := new(models.Category)
	p.ViewArgs["categorys"] = categoryModel.FindAll()
	tags, err := tagModel.ListAll()
	createtime := time.Now()
	if err != nil {
		tags = make([]models.Tag, 0)
	}
	blog := &models.Blog{Id: postid}
	if postid > 0 {
		blog, err = blog.FindById()
		if err != nil {
			return p.NotFound("博客不存在")
		}
		createtime = blog.CreateTime
	}
	p.ViewArgs["blog"] = blog
	p.ViewArgs["tags"] = tags
	p.ViewArgs["createtime"] = createtime
	return p.RenderTemplate("Admin/Post/Index.html")
}

//ManagePost .
// 管理博客页面
func (p *Post) ManagePost(uid, category int64) revel.Result {
	blogs, err := blogModel.GetBlogByPageAND(uid, category, 1, 20)
	if err != nil {
		blogs = make([]models.Blog, 0)
	}
	p.ViewArgs["blogs"] = blogs
	p.ViewArgs["p_uid"] = uid
	p.ViewArgs["p_ca"] = category
	return p.RenderTemplate("Admin/Post/Manage-post.html")
}

// NewPostHandler to Add new article.
// 添加博客
func (p *Post) NewPostHandler() revel.Result {
	data := new(PostData)
	p.Params.Bind(&data, "data")
	p.Validation.Required(data.Title).Message("标题不能为空")
	p.Validation.Required(data.ContentHTML).Message("内容不能为空")

	if p.Validation.HasErrors() {
		return p.RenderJSON(&ResultJson{Success: false, Msg: p.Validation.Errors[0].Message})
	}

	blog := new(models.Blog)
	blog.Title = data.Title
	if data.Ident == "" {
		data.Ident = blog.Title
	}
	blog.Ident = data.Ident
	blog.ContentHTML = data.ContentHTML
	blog.ContentMD = data.ContentMD
	blog.CategoryId = data.Category
	blog.Type = data.Type
	blog.Summary = data.Summary

	// 处理创建时间
	tm, err := time.Parse("2006-01-02", data.Createtime)
	if err != nil {
		blog.CreateTime = time.Now()
	} else {
		blog.CreateTime = tm
	}

	uid := p.Session["UID"]
	authorid, _ := strconv.Atoi(uid)
	blog.CreateBy = int64(authorid)

	if data.passwd != "" {
		blog.Passwd = data.passwd
	}

	var blogID int64
	if data.Id > 0 {
		blog.Id = data.Id
		_, err = blog.Update()
		if err != nil {
			revel.ERROR.Println("博客更新失败：", err)
		}
		blogID = data.Id
	} else {
		blogID, err = blog.New()
	}

	// 删除所有的旧标签
	blog.DeleteAllBlogTags()
	// 添加新的标签
	btr := new(models.BlogTag)
	newTags := strings.Split(data.NewTag, ",")
	for _, v := range newTags {
		tagid, err := tagModel.NewTagByName(v)
		if tagid > 0 {
			btr.AddTagRef(tagid, blogID)
		} else {
			revel.ERROR.Println("创建标签失败：", err)
		}
	}

	// 处理标签关联
	tagids := strings.Split(data.Tag, ",")
	for _, v := range tagids {
		id, err := strconv.Atoi(v)
		if err == nil {
			btr.AddTagRef(int64(id), blogID)
		}
	}

	if err != nil || blogID <= 0 {
		p.Flash.Error("msg", "create new blogger post error.")
		return p.RenderJSON(&ResultJson{Success: false, Msg: err.Error(), Data: ""})
	}
	return p.RenderJSON(&ResultJson{Success: true})
}

func (p *Post) QueryCategorys() revel.Result {
	c := new(models.Category)
	arr := c.FindAll()
	return p.RenderJSON(&ResultJson{Success: true, Msg: "", Data: arr})
}

// CreateTag to create a new tag when create a blog
// 在创建博客的时候创建一个标签
func (p *Post) CreateTag(name string) revel.Result {
	_, err := tagModel.NewTagByName(name)
	if err != nil {
		revel.ERROR.Println("创建标签失败：", err)
		return p.RenderJSON(&ResultJson{Success: false, Msg: err.Error(), Data: ""})
	}
	return p.RenderJSON(&ResultJson{Success: true, Msg: "", Data: ""})
}

// Delete a blog
// 删除博客
func (p *Post) Delete(ids string) revel.Result {
	idArr := strings.Split(ids, ",")
	if len(idArr) <= 0 {
		return p.RenderJSON(&ResultJson{Success: false, Msg: "参数无效"})
	}
	validIdArr := []int64{}
	for _, v := range idArr {
		id, err := strconv.Atoi(v)
		if err == nil {
			validIdArr = append(validIdArr, int64(id))
		}
	}
	ok, err := blogModel.BatchDel(validIdArr)
	if ok {
		return p.RenderJSON(&ResultJson{Success: true})
	}
	return p.RenderJSON(&ResultJson{Success: false, Msg: err})
}
