package controllers

import (
	"blog/app/models"
	"blog/app/routes"
	"blog/app/support"
	"encoding/json"
	"strconv"
	"time"

	"github.com/revel/revel"
)

//Login controller
type Login struct {
	*revel.Controller
}

//SignIn page.
func (l Login) SignIn() revel.Result {
	return l.Render()
}

//SignInHandler -> handle Sign
func (l Login) SignInHandler(name, passwd string) revel.Result {

	model := &models.Admin{Name: name, Passwd: passwd}
	admin, err := model.SignIn(l.Request)

	if err != "" {
		revel.ERROR.Println(err)
		l.Flash.Error("msg", err)
		return l.Redirect(routes.Login.SignIn())
	}

	revel.INFO.Println(admin)
	//put admin id in seesion
	l.Session["UID"] = strconv.Itoa(admin.Id)
	//set admin info in cache, time out time.Minute * 30
	data, _ := json.Marshal(&admin)
	support.Cache.Set(support.SPY_ADMIN_INFO+strconv.Itoa(admin.Id), string(data), time.Minute*30)

	l.Flash.Success("msg", "success")
	return l.RenderHtml(l.Session.Id())
}

//SignUp page.
func (l Login) SignUp() revel.Result {
	return l.Render()
}

//SignUpHandler -> handle sign up.
func (l Login) SignUpHandler(name, email, passwd string) revel.Result {

	model := &models.Admin{Name: name, Email: email, Passwd: passwd}
	has, err := model.New()

	if err != "" && has <= 0 {
		revel.ERROR.Println(err)
		l.Flash.Error("msg", err)
		return l.Redirect(routes.Login.SignUp())
	}

	return l.RenderHtml("ok")
}
