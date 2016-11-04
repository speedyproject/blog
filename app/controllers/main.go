package controllers

import "github.com/revel/revel"

//Main Controller.
type Main struct {
	*revel.Controller
}

//Main page.
func (m *Main) Main() revel.Result {

	return m.Render()
}
