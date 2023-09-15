package repository

import (
	"strconv"

	config "github.com/Hareshutit/ShopEase/config/auth"
	"github.com/redis/go-redis/v9"
)

func CreateRedisRepository(cfg config.Config) RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host + ":" + strconv.Itoa(cfg.Redis.Port),
		DB:   0, // use default DB
	})

	return RedisRepository{client}
}
