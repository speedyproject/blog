package controllers

import (
	"blog/app/support"
	"fmt"

	"github.com/revel/revel"
)

type DBParams struct {
	Db_host string
	Db_user string
	Db_pass string
	Db_port int
	Db_name string
	Driver  string
}

type AdminParams struct {
	Admin_user  string
	Admin_pass  string
	Admin_email string
}

type Install struct {
	*revel.Controller
}

func (i *Install) Index() revel.Result {
	return i.Render()
}

func (i *Install) HandleInstall() revel.Result {
	return nil
}

func (i *Install) AddAdmin() revel.Result {
	params := new(AdminParams)
	i.Params.Bind(params, "info")
	var err error
	// err := (params)

	if err != nil {
		return i.RenderJson(&ResultJson{Success: false, Msg: err.Error(), Data: ""})
	}
	return i.RenderJson(&ResultJson{Success: true})
}

func (i *Install) AddDB() revel.Result {
	params := new(DBParams)
	i.Params.Bind(params, "info")
	params.Driver = "mysql"
	err := i.checkDB(params)
	if err != nil {
		return i.RenderJson(&ResultJson{Success: false, Msg: err.Error(), Data: ""})
	}
	return i.RenderJson(&ResultJson{Success: true})
}

func (i *Install) checkDB(info *DBParams) error {
	fmt.Println("db info: ", info)
	err := support.TestXorm(info.Driver, info.Db_user, info.Db_pass, info.Db_host, info.Db_name, info.Db_port)
	if err != nil {
		return err
	}
	err = support.AddDB(info.Db_host, fmt.Sprintf("%d", info.Db_port), info.Db_user, info.Db_pass, info.Db_name, info.Driver)
	if err != nil {
		return err
	}
	return nil
}
