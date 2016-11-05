package controllers

import (
	"blog/app/models"

	"github.com/revel/revel"
)

// User for User Controller
type User struct {
	Admin
}

// Main to list all users
func (user *User) Main() revel.Result {
	admin := models.Admin{}
	users, err := admin.List()
	if err != nil {
		revel.ERROR.Println(err)
		return nil
	}
	user.RenderArgs["users"] = users
	return user.RenderTemplate("Admin/User/Main.html")
}

// Edit User
func (user *User) Edit() revel.Result {
	return nil
}

// Create user
func (user *User) Create() revel.Result {
	return nil
}

// Update user info
func (user *User) Update() revel.Result {
	return nil
}

// Delete a user
func (user *User) Delete() revel.Result {
	return nil
}
