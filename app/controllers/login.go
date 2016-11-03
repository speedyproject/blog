package controllers

import "github.com/revel/revel"

type Login struct {
	*revel.Controller
}

func (l Login) SignIn() revel.Result {
	return l.Render()
}


func (l Login) SignInHandler() revel.Result {
	return l.RenderHtml("ok")
}