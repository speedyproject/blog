package support

import "github.com/revel/revel/cache"

var cacheType int
var MCache cache.Cache
var SPY_CONF_MD5_VAL string
var SPY_CONF_SIGN_VAL string

const (
	DEFAULT = 0
	REDIS   = 1
)

func InitCache() {
	// Cache := cache
	MCache = cache.NewInMemoryCache(cache.DEFAULT)
	InitRedis()
	loadCache()
}

//Load config data to redis cache.
func loadCache() {
	md5_key, _ := AppConfig.String("secret", "secret.md5.key")
	sign_key, _ := AppConfig.String("secret", "secret.sign.key")

	MCache.Set(SPY_CONF_MD5_KEY, md5_key, cache.FOREVER)
	MCache.Set(SPY_CONF_SIGN_KEY, sign_key, cache.FOREVER)
	SPY_CONF_MD5_VAL = md5_key
	SPY_CONF_SIGN_VAL = sign_key
}
