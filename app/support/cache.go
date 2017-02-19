package support

import (
	"math/rand"
	"time"

	"github.com/revel/config"
	"github.com/revel/revel/cache"
)

var cacheType int
var MCache cache.Cache
var SPY_CONF_MD5_VAL string
var SPY_CONF_SIGN_VAL string
var AppConfig *config.Config
var IsInstalled bool

const (
	DEFAULT           = 0
	REDIS             = 1
	SPY_CONF_MD5_KEY  = "speedy:conf:md5:key"
	SPY_CONF_SIGN_KEY = "speedy:conf:sign:key"

	SPY_ADMIN_INFO = "admin:info:id:"

	SPY_BLOGGER_LIST   = "speedy:blogger:list"
	SPY_BLOGGER_SINGLE = "speedy:blogger:id:"
)

func InitCache(isInstalled bool, config *config.Config) {
	// Cache := cache
	AppConfig = config
	IsInstalled = isInstalled
	MCache = cache.NewInMemoryCache(cache.DEFAULT)
	InitRedis()
	loadCache(isInstalled)
}

//Load config data to redis cache.
func loadCache(hasConfig bool) {
	var md5_key, sign_key string
	if hasConfig {
		md5_key, _ = AppConfig.String("secret", "secret.md5.key")
		sign_key, _ = AppConfig.String("secret", "secret.sign.key")
	} else {
		md5_key = randStringBytes(16)
		sign_key = randStringBytes(16)
	}

	MCache.Set(SPY_CONF_MD5_KEY, md5_key, cache.FOREVER)
	MCache.Set(SPY_CONF_SIGN_KEY, sign_key, cache.FOREVER)
	AppConfig.AddOption("secret", "secret.md5.key", md5_key)
	AppConfig.AddOption("secret", "secret.sign.key", sign_key)
	SPY_CONF_MD5_VAL = md5_key
	SPY_CONF_SIGN_VAL = sign_key
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
