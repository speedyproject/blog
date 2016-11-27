package support

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/revel/config"
	"github.com/revel/revel"
)

var Xorm *xorm.Engine

//Init the mysql.
func InitXorm() {

	file := (revel.BasePath + "/conf/speedy.conf")
	data, _ := config.ReadDefault(file)

	driver, _ := data.String("database", "database.driver")
	dbname, _ := data.String("database", "database.dbname")
	user, _ := data.String("database", "database.user")
	passwd, _ := data.String("database", "database.password")
	host, _ := data.String("database", "database.host")
	prefix, _ := data.String("database", "database.prefix")

	var err error
	Xorm, err = xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, passwd, host, dbname))
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
	Xorm.SetTableMapper(tbMapper)
	Xorm.ShowSQL(true)
	if err != nil {
		revel.ERROR.Println(err)
	} else {
		Xorm.Ping()
	}
}

func TestXorm(driver, user, pass, host, dbname string, port int) error {
	var err error
	Xorm, err = xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pass, host, port, dbname))
	if err != nil {
		return err
	}
	return Xorm.Ping()
}
