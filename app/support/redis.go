package support

import (
	"github.com/alecthomas/log4go"
	"github.com/revel/config"
	"github.com/revel/revel"
	"gopkg.in/redis.v5"
)

var Cache *redis.Client

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
