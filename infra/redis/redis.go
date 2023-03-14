package redis

import (
	"chalet/pkg/entity"
	"log"

	"github.com/go-redis/redis"
)

var client = &redis.Client{}

func Init(redisConfig entity.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("init redis failed:", err)
	}
}
