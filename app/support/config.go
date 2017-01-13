package support

import (
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
var IsInstalled bool

func InitConfig() {
	file := (revel.BasePath + CONFIG_PATH)
	var err error
	//检查配置文件是否存在
	AppConfig, err = config.ReadDefault(file)
	IsInstalled = true
	if err != nil {
		revel.WARN.Println("获取配置文件失败，准备安装", err)
		IsInstalled = false
	} else {
		// 检查数据库是否可以正常连接
		err = initXorm()
		if err != nil {
			IsInstalled = false
			revel.WARN.Println("连接数据库失败，准备安装")
		}
		// FIXME：Laily 检查表是否存在，然后检查是否同步表结构
	}
	if !IsInstalled {
		AppConfig = config.New(config.DEFAULT_COMMENT, config.ALTERNATIVE_SEPARATOR, false, true)
	} else {
		revel.INFO.Println("加载配置文件成功...")
	}
	InitCache()
}

func AddDB(dbhost, dbport, dbuser, dbpass, dbname, dbtype string) error {
	AppConfig.AddOption("database", "database.host", dbhost)
	AppConfig.AddOption("database", "database.port", dbport)
	AppConfig.AddOption("database", "database.user", dbuser)
	AppConfig.AddOption("database", "database.password", dbpass)
	AppConfig.AddOption("database", "database.dbname", dbname)
	AppConfig.AddOption("database", "database.driver", dbtype)
	return nil
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

func FinishInstall() {
	err := writeConfig()
	if err != nil {
		revel.ERROR.Println("write config error: ", err)
	}
}
