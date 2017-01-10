package support

import (
	"log"
	"os"

	"github.com/revel/config"
	"github.com/revel/revel"
)

const (
	SPY_CONF_MD5_KEY  = "speedy:conf:md5:key"
	SPY_CONF_SIGN_KEY = "speedy:conf:sign:key"

	SPY_ADMIN_INFO = "admin:info:id:"

	SPY_BLOGGER_LIST   = "speedy:blogger:list"
	SPY_BLOGGER_SINGLE = "speedy:blogger:id:"

	CONFIG_PATH = "/conf/speedy.conf"
)

var AppConfig *config.Config
var isInstalled bool

func InitConfig() {
	file := (revel.BasePath + CONFIG_PATH)
	var err error
	AppConfig, err = config.ReadDefault(file)
	if err != nil {
		log.Println("log config error: ", err)
		AppConfig = config.New(config.DEFAULT_COMMENT, config.ALTERNATIVE_SEPARATOR, false, true)
		// AppConfig.AddOption("appconfig", "installed", "false")
		isInstalled = false
	}
}

func AddDB(dbhost, dbport, dbuser, dbpass, dbname, dbtype string) error {
	AppConfig.AddOption("dbconfig", "dbhost", dbhost)
	AppConfig.AddOption("dbconfig", "dbport", dbport)
	AppConfig.AddOption("dbconfig", "dbuser", dbuser)
	AppConfig.AddOption("dbconfig", "dbpass", dbpass)
	AppConfig.AddOption("dbconfig", "dbname", dbname)
	AppConfig.AddOption("dbconfig", "dbtype", dbtype)

	err := writeConfig()
	return err
}

func writeConfig() error {
	filepath := revel.BasePath + CONFIG_PATH
	_, err := os.Open(filepath)
	if err != nil {
		os.Create(filepath)
	}

	err = AppConfig.WriteFile(filepath, 0775, "default config")
	if err != nil {
		return err
	}
	return nil
}
