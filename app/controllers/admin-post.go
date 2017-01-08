package controllers

import (
	"blog/app/models"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/revel/revel"
)

/**
 * Add a blog for admin user
 * 发布博客 action
 */

// PostData model.
// 发布博客前端提交的数据
type PostData struct {
	Title       string //博客标题
	ContentMD   string //博客内容 MD
	ContentHTML string // 博客内容 HTML
	Category    int    // 博客类别
	Tag         string // 标签 格式：12,14,32
	Keywords    string // 关键词 格式：java,web开发
	passwd      string //博客内容是否加密
	Summary     string // 博客摘要
	Type        int    // 0 表示 markdown，1 表示 html
}

// User for User Controller
type Post struct {
	Admin
}

// 创建博客页面
func (p *Post) Index() revel.Result {
	categoryModel := new(models.Category)
	tagModel := new(models.BloggerTag)
	p.RenderArgs["categorys"] = categoryModel.FindAll()
	tags, err := tagModel.ListAll()
	if err != nil {
		tags = make([]models.BloggerTag, 0)
	}
	p.RenderArgs["tags"] = tags
	p.RenderArgs["timenow"] = time.Now()
	return p.RenderTemplate("Admin/Post/Index.html")
}

//ManagePost .
// 管理博客页面
func (p *Post) ManagePost(uid, category int64) revel.Result {
	blogModel := new(models.Blogger)
	blogs, err := blogModel.GetBlogByPageAND(uid, category, 1, 20)
	if err != nil {
		blogs = make([]models.Blogger, 0)
	}
	p.RenderArgs["blogs"] = blogs
	p.RenderArgs["p_uid"] = uid
	p.RenderArgs["p_ca"] = category
	return p.RenderTemplate("Admin/Post/Manage-post.html")
}

// NewPostHandler to Add new article.
// 添加博客
func (p *Post) NewPostHandler() revel.Result {
	data := new(PostData)
	p.Params.Bind(&data, "data")
	fmt.Println("data= ", data)
	p.Validation.Required(data.Title).Message("title can't be null.")
	p.Validation.Required(data.ContentHTML).Message("context can't be null.")

	if p.Validation.HasErrors() {
		p.Validation.Keep()
		p.FlashParams()
		// TODO Redirect new post page.
	}

	blog := new(models.Blogger)
	blog.Title = data.Title
	blog.ContentHTML = data.ContentHTML
	blog.ContentMD = data.ContentMD
	blog.CategoryId = data.Category
	blog.Type = data.Type
	blog.Summary = data.Summary
	uid := p.Session["UID"]
	id, _ := strconv.Atoi(uid)

	blog.CreateBy = id

	if data.passwd != "" {
		blog.Passwd = data.passwd
	}

	has, err := blog.New()

	if err != nil || has <= 0 {
		p.Flash.Error("msg", "create new blogger post error.")
		return p.RenderJson(&ResultJson{Success: false, Msg: err.Error(), Data: ""})
		// TODO Redirect new post page.
	}
	return p.RenderHtml("ok")
}

func (p *Post) QueryCategorys() revel.Result {
	c := new(models.Category)
	arr := c.FindAll()
	return p.RenderJson(&ResultJson{Success: true, Msg: "", Data: arr})
}

func (p *Post) CreateTag(name string) revel.Result {
	tag := new(models.BloggerTag)
	tag.Name = name
	tag.Parent = 0
	tag.Type = 0
	_, err := tag.New()
	if err != nil {
		return p.RenderJson(&ResultJson{Success: false, Msg: err.Error(), Data: ""})
	}
	return p.RenderJson(&ResultJson{Success: true, Msg: "", Data: ""})
}

func (p *Post) Delete(ids string) revel.Result {
	idArr := strings.Split(ids, ",")
	if len(idArr) > 0 {
		for _, v := range idArr {
			id, err := strconv.Atoi(v)
			if err == nil {
				blog := &models.Blogger{Id: int64(id)}
				blog.Del()
			}
		}
		return p.RenderJson(&ResultJson{Success: true})
	} else {
		return p.RenderJson(&ResultJson{Success: true})
	}
}
