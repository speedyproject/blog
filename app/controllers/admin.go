package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"blog/app/models"

	"github.com/revel/revel"
)

//Admin controller.
type Admin struct {
	*revel.Controller
}

// AdminChecker for get the value of which module user choose,
// and make the menu selected.
// TODO:Laily check if it is a admin user
func (admin *Admin) AdminChecker() revel.Result {
	url := fmt.Sprintf("%s", admin.Request.URL.Path)
	revel.INFO.Println(url)
	uriStr := strings.Split(url, "/")
	if len(uriStr) > 0 {
		uri := uriStr[len(uriStr)-1]
		switch len(uriStr) {
		case 4:
			uri = uriStr[len(uriStr)-2]
		case 5:
			uri = uriStr[len(uriStr)-3]
		case 6:
			uri = uriStr[len(uriStr)-4]
		}
		if uri == "" || uri == "main" {
			admin.RenderArgs["managementPage"] = "index"
		} else {
			admin.RenderArgs["managementPage"] = uri
		}
		revel.INFO.Printf("Admin::AdminChecker uri -> %s", uri)
	}
	return nil
}

//Main page.
func (a *Admin) Main() revel.Result {
	return a.Render()
}

//PostData model.
type PostData struct {
	Title    string
	Context  string
	Date     time.Time
	Label    string
	Tag      string
	Keywords string
	passwd   string
}

// Add new article.
func (a *Admin) NewArticleHandler() revel.Result {

	data := new(PostData)
	a.Params.Bind(&data, "data")

	a.Validation.Required(data.Title).Message("title can't be null.")
	a.Validation.Required(data.Context).Message("context can't be null.")
	a.Validation.Required(data.Date).Message("date can't be null.")

	if a.Validation.HasErrors() {
		a.Validation.Keep()
		a.FlashParams()
		// TODO Redirect new post page.
	}

	blog := new(models.Blogger)
	blog.Title = data.Title
	blog.Context = data.Context
	blog.CreateTime = data.Date

	uid := a.Session["UID"]
	id, _ := strconv.Atoi(uid)

	blog.CreateBy = id

	if data.passwd != "" {
		blog.Passwd = data.passwd
	}

	has, err := blog.New()

	if err != nil || !has {
		a.Flash.Error("msg", "create new blogger post error.")
		// TODO Redirect new post page.
	}

	return a.RenderHtml("ok")
}
