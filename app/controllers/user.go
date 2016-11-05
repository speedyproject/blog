package controllers

import "github.com/revel/revel"

// User for User Controller
type User struct {
	*revel.Controller
}

// Index to list all users
func (user *User) Main() revel.Result {
	return user.Render()
}
