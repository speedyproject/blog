package controllers

import (
	"github.com/revel/revel"
)

// Blogger controller
type Blogger struct {
	*revel.Controller
}

//Blogger page.
func (b Blogger) BloggerPage() revel.Result {
	return b.Render()
}
