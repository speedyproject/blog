package controllers

import (
	"fmt"
	"strings"

	"github.com/revel/revel"
)

//Admin controller.
type Admin struct {
	*revel.Controller
}

// AdminChecker for get the value of which module user choose,
// and make the menu selected.
func (admin *Admin) AdminChecker() revel.Result {
	url := fmt.Sprintf("%s", admin.Request.URL.Path)
	if strings.Contains(url, "admin/user") {
		admin.RenderArgs["managementPage"] = "user"
	} else {
		admin.RenderArgs["managementPage"] = "index"
	}
	revel.INFO.Println(url)
	return nil
}

//Main page.
func (a *Admin) Main() revel.Result {
	return a.Render()
}
