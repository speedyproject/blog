package support

import (
	"github.com/revel/config"
	"github.com/revel/revel"
)

const (
	SPY_CONF_MD5_KEY  = "speedy:conf:md5:key"
	SPY_CONF_SIGN_KEY = "speedy:conf:sign:key"

	SPY_ADMIN_INFO = "admin:info:id:"

	SPY_BLOGGER_LIST   = "speedy:blogger:list"
	SPY_BLOGGER_SINGLE = "speedy:blogger:id:"
)

var AppConfig *config.Config

func InitConfig() {
	file := (revel.BasePath + "/conf/speedy.conf")
	var err error
	AppConfig, err = config.ReadDefault(file)
	if err != nil {
		revel.INFO.Println("log config error: ", err)
	}
}
