package support

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

var Xorm *xorm.Engine

//Init the Xorm.
func initXorm() error {
	dbdriver, _ := AppConfig.String("database", "database.driver")
	switch dbdriver {
	case "mysql":
		return initMySQL()
	}
	return errors.New("no db driver")
}

func initMySQL() error {
	dbname, _ := AppConfig.String("database", "database.dbname")
	user, _ := AppConfig.String("database", "database.user")
	passwd, _ := AppConfig.String("database", "database.password")
	host, _ := AppConfig.String("database", "database.host")
	port, _ := AppConfig.String("database", "database.port")
	prefix, _ := AppConfig.String("database", "database.prefix")

	var err error
	Xorm, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, passwd, host, port, dbname))
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
	Xorm.SetTableMapper(tbMapper)
	Xorm.ShowSQL(true)
	if err != nil {
		revel.ERROR.Println(err)
		return err
	}
	return Xorm.Ping()
}

func TestXorm(driver, user, pass, host, dbname string, port int, prefix string) error {
	var err error
	Xorm, err = xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pass, host, port, dbname))
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
	Xorm.SetTableMapper(tbMapper)
	Xorm.ShowSQL(true)
	if err != nil {
		return err
	}
	return Xorm.Ping()
}
