package controllers

import (
	"blog/app/support"

	"github.com/revel/revel"
)

type InstallParams struct {
	Db_host     string
	Db_user     string
	Db_pass     string
	Db_port     int
	Db_name     string
	Admin_user  string
	Admin_pass  string
	Admin_email string
	Driver      string
}

type Install struct {
	*revel.Controller
}

func (i *Install) Index() revel.Result {
	return i.Render()
}

func (i *Install) HandleInstall() revel.Result {
	params := new(InstallParams)
	i.Params.Bind(params, "info")
	params.Driver = "mysql"
	err := i.checkDB(params)
	if err != nil {
		return i.RenderJson(&ResultJson{Success: false, Msg: err.Error(), Data: "db"})
	}
	return nil
}

func (i *Install) checkDB(info *InstallParams) error {
	err := support.TestXorm(info.Driver, info.Db_user, info.Db_pass, info.Db_host, info.Db_name, info.Db_port)
	if err != nil {
		return err
	}
	return nil
}
