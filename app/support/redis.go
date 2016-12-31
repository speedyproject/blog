package support

import (
	"github.com/revel/revel"
	"gopkg.in/redis.v5"
)

var Cache *redis.Client

//Init the redis client.
func InitRedis() {
	data := AppConfig
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
		revel.ERROR.Println(err)
	} else {
		revel.INFO.Println(res)
	}
}
