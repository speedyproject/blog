package controllers

import (
	"blog/app/models"
	"strconv"

	"github.com/revel/revel"
)

// BlogTag controller
type BlogTag struct {
	*revel.Controller
}

//
func (b *BlogTag) Index(ident string) revel.Result {
	tag := new(models.BloggerTag)
	tag, err := tag.GetByIdent(ident)
	if err != nil {
		revel.ERROR.Panic("wrong")
	}
	blogs := tag.FindBlogByTag("")
	b.RenderArgs["flag"] = "tag"
	b.RenderArgs["tag"] = tag
	b.RenderArgs["blogs"] = blogs
	return b.RenderTemplate("Main/Blog4Search.html")
}

// GetAllTags to find all tags
// 获取所有的标签
func (b *BlogTag) GetAllTags() revel.Result {
	tagModel := new(models.BloggerTag)
	tags, err := tagModel.ListAll()
	if err != nil {
		revel.ERROR.Println("find all tags error: ", err)
	}
	return b.RenderJson(tags)
}

// QueryTags to Search for tag
// 根据用户输入的单词匹配 tag
func (b *BlogTag) QueryTags(t string) revel.Result {
	tag := new(models.BloggerTag)
	res, err := tag.QueryTags(t)
	if err != nil {
		return b.RenderJson(&ResultJson{Success: false, Msg: err.Error(), Data: ""})
	}
	resMap := make(map[int64]string, 0)
	for _, v := range res {
		id, err := strconv.Atoi(string(v["id"]))
		if err == nil {
			resMap[int64(id)] = string(v["name"])
		}
	}
	return b.RenderJson(&ResultJson{Success: true, Msg: "", Data: resMap})
}
