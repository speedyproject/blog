package controllers

import "github.com/revel/revel"

type Login struct {
	*revel.Controller
}

func (l Login) SignIn() revel.Result {

	return l.Render()
}