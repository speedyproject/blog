package support

import (
	"github.com/alecthomas/log4go"
	"github.com/revel/config"
	"github.com/revel/revel"
	"gopkg.in/redis.v5"
)

var Cache *redis.Client

const (
	SPY_CONF_MD5_KEY  = "speedy:conf:md5:key"
	SPY_CONF_SIGN_KEY = "speedy:conf:sign:key"

	USER_LOGIN_BY_SIGN = "admin:user:sign:"
	USER_DATA_BY_ID    = "admin:user:id:"
)

//Init the redis client.
func InitRedis() {

	file := (revel.BasePath + "/conf/speedy.conf")
	data, _ := config.ReadDefault(file)

	host, _ := data.String("redis", "redis.host")
	passwd, _ := data.String("redis", "redis.password")
	poolsize, _ := data.Int("redis", "redis.poolsize")

	Cache = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: passwd,
		DB:       0,
		PoolSize: poolsize,
	})

	res, err := Cache.Ping().Result()

	if err != nil {
		log4go.Error(err)
	} else {
		log4go.Debug(res)
	}
}

//Load config data to redis cache.
func LoadCache() {
	file := (revel.BasePath + "/conf/speedy.conf")
	data, _ := config.ReadDefault(file)

	md5_key, _ := data.String("secret", "secret.md5.key")
	sign_key, _ := data.String("secret", "secret.sign.key")

	Cache.Set(SPY_CONF_MD5_KEY, md5_key, 0)
	Cache.Set(SPY_CONF_SIGN_KEY, sign_key, 0)
}
