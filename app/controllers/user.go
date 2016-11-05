package controllers

import "github.com/revel/revel"

// User for User Controller
type User struct {
	Admin
}

// Main to list all users
func (user *User) Main() revel.Result {
	return user.RenderTemplate("Admin/User/Main.html")
}
