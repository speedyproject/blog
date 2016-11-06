package controllers

import (
	"blog/app/models"
	"strings"

	"strconv"

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
func (user *User) Edit(id int64) revel.Result {
	u := &models.Admin{}
	u, err := u.GetUserByID(id)
	if err != nil {
		return user.RenderText("user is not exist")
	}
	user.RenderArgs["user"] = u
	return user.RenderTemplate("Admin/User/Edit.html")
}

func (user *User) EditHandler(username, nickname, password, email string, group int, id int64) revel.Result {
	u := &models.Admin{Name: username, Nickname: nickname, Email: email, Passwd: password, RoleId: group}
	_, err := u.UpdateAdmin(id, u)
	if err != "" {
		return user.RenderJson(ResultJson{Success: false, Msg: err})
	}
	return user.RenderJson(ResultJson{Success: true, Data: id})
}

// Create user
func (user *User) Create() revel.Result {
	return user.RenderTemplate("Admin/User/Create.html")
}

// CreateHandler for do create user
func (user *User) CreateHandler(username, nickname, password, email string, group int) revel.Result {
	revel.INFO.Printf("%s,%s,%s,%s,%d", username, nickname, password, email, group)
	u := &models.Admin{Name: username, Nickname: nickname, Email: email, Passwd: password, RoleId: group}
	id, err := u.New()
	if err != "" {
		return user.RenderJson(ResultJson{Success: false, Msg: err})
	}
	return user.RenderJson(ResultJson{Success: true, Data: id})
}

// Delete a user
func (user *User) Delete(ids string) revel.Result {
	arr := strings.Split(ids, ",")
	u := new(models.Admin)
	for _, v := range arr {
		id, err := strconv.Atoi(v)
		if err == nil {
			u.DeleteAdmin(int64(id))
		}
	}
	return nil
}
