package controllers

import (
	"fmt"
	"strings"
	"time"

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
	Content  string
	Date     time.Time
	Label    string
	Tag      string
	Keywords string
	passwd   string
}
