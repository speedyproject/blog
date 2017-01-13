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
	loadCache(IsInstalled)
}

//Load config data to redis cache.
func loadCache(hasConfig bool) {
	var md5_key, sign_key string
	if hasConfig {
		md5_key, _ = AppConfig.String("secret", "secret.md5.key")
		sign_key, _ = AppConfig.String("secret", "secret.sign.key")
	} else {
		md5_key = "yigeheshangtiaoshuihe"
		sign_key = "lianggeheshangtaishuihe"
	}

	MCache.Set(SPY_CONF_MD5_KEY, md5_key, cache.FOREVER)
	MCache.Set(SPY_CONF_SIGN_KEY, sign_key, cache.FOREVER)
	AppConfig.AddOption("secret", "secret.md5.key", md5_key)
	AppConfig.AddOption("secret", "secret.sign.key", sign_key)
	SPY_CONF_MD5_VAL = md5_key
	SPY_CONF_SIGN_VAL = sign_key
}
