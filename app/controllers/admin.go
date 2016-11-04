package controllers

import (
	"github.com/revel/revel"
)

//Admin controller.
type Admin struct {
	*revel.Controller
}

//Main page.
func (a *Admin) Main() revel.Result {
	return a.Render()
}
