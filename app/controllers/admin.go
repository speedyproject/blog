package controllers

import (
	"fmt"
	"strings"

	"github.com/revel/revel"
)

//Admin controller.
// Admin 控制器，后代所有操作都在此控制器下
type Admin struct {
	*revel.Controller
}

// AdminChecker for get the value of which module user choose,
// and make the menu selected.
// TODO:Laily check if it is a admin user
// 检测 url 是访问后台哪个模块的操作
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
			admin.ViewArgs["managementPage"] = "index"
		} else {
			admin.ViewArgs["managementPage"] = uri
		}
	}
	return nil
}

//Main page.
// 后台首页
func (a *Admin) Main() revel.Result {
	return a.Render()
}
