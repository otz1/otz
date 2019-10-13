package cache

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

// RedisCache is a wrapper of the redis client
type RedisCache struct {
	*redis.Client
}

// RedisDAO global singleton instance of the redis cache
// which is for use internally in this package.
var redisDAO = newRedisCache()

func newRedisCache() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0, // use default DB
	})
	return &RedisCache{client}
}

// Close will close the redis connection if one exists
func Close() {
	if redisDAO != nil {
		err := redisDAO.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
