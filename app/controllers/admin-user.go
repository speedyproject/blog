package controllers

import (
	"blog/app/models"
	"strconv"
	"strings"

	"github.com/revel/revel"
)

// User for User Controller
// 用户管理控制器
type User struct {
	Admin
}

// Main to list all users
// 用户管理首页，列出所有用户
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
// 编辑用户
func (user *User) Edit(id int64) revel.Result {
	u := &models.Admin{}
	u, err := u.GetUserByID(id)
	if err != nil {
		return user.RenderText("user is not exist")
	}
	user.RenderArgs["user"] = u
	return user.RenderTemplate("Admin/User/Edit.html")
}

// EditHandler to update user info
// 处理编辑用户的方法
func (user *User) EditHandler(username, nickname, password, email string, group int, id int64) revel.Result {
	u := &models.Admin{Name: username, Nickname: nickname, Email: email, Passwd: password, RoleId: int64(group)}
	_, err := u.UpdateAdmin(id, u)
	if err != "" {
		return user.RenderJson(ResultJson{Success: false, Msg: err})
	}
	return user.RenderJson(ResultJson{Success: true, Data: id})
}

// Create user
// 创建用户页面
func (user *User) Create() revel.Result {
	return user.RenderTemplate("Admin/User/Create.html")
}

// CreateHandler for do create user
// 处理创建用户的请求
func (user *User) CreateHandler(username, nickname, password, email string, group int) revel.Result {
	revel.INFO.Printf("%s,%s,%s,%s,%d", username, nickname, password, email, group)
	u := &models.Admin{Name: username, Nickname: nickname, Email: email, Passwd: password, RoleId: int64(group)}
	id, err := u.New()
	if err != "" {
		return user.RenderJson(ResultJson{Success: false, Msg: err})
	}
	return user.RenderJson(ResultJson{Success: true, Data: id})
}

// Delete a user
// 删除用户
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
