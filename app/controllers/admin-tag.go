package controllers

import (
	"blog/app/models"
	"log"
	"github.com/revel/revel"
	"strings"
	"strconv"
)

/**
 * 后台管理 tag 的控制器
 */

var tagModel = new(models.BloggerTag)

type AdminTag struct {
	Admin
}

// admin manage tag index page
// 后台管理标签首页
func (a *AdminTag) Index() revel.Result {
	tagModel := new(models.BloggerTag)
	list, err := tagModel.ListAll()
	if err != nil {
		log.Panic("in admin-tag page index error: ", err)
	}
	a.RenderArgs["tags"] = list
	return a.RenderTemplate("Admin/Tag/Index.html")
}

// 编辑标签
func (a *AdminTag) Edit(tagID int64, tagName, tagIdent string) revel.Result {
	a.Validation.Required(tagID)
	a.Validation.Required(tagName)
	a.Validation.Required(tagIdent)
	if a.Validation.HasErrors() {
		return a.RenderJson(&ResultJson{Success:false, Msg:a.Validation.Errors[0].Message, Data:""})
	}
	tag := new(models.BloggerTag)
	tag, err := tag.GetByID(tagID)
	if err != nil {
		return a.RenderJson(&ResultJson{Success:false, Msg:err.Error(), Data:""})
	}
	tag.Ident = tagIdent
	tag.Name = tagName
	if !tag.Update() {
		return a.RenderJson(&ResultJson{Success:false, Msg:"更新失败", Data:""})
	}
	return a.RenderJson(&ResultJson{Success:true, Msg:"", Data:""})
}

// 删除标签
func (a *AdminTag) Del(ids string) revel.Result {
	idsArr := strings.Split(ids, ",")
	if len(idsArr) > 0 {
		for i, v := range idsArr {
			_, err := strconv.Atoi(v)
			if err != nil {
				idsArr = append(idsArr[:i], idsArr[i + 1:]...)
			}
		}
		tagModel.Delete(idsArr)
	}
	return a.RenderJson(&ResultJson{Success:true, Msg:"", Data:""})
}
