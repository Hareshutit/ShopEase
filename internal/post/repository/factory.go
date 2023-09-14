package repository

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/Hareshutit/ShopEase/config/post"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgx/stdlib"
)

func newPool(cfg config.Config) *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "redis_cart:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func CreatePostgressRepository(cfg config.Config) PostPostgressRepository {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Db.User, cfg.Db.DataBaseName, cfg.Db.Password, cfg.Db.Host,
		cfg.Db.Port, cfg.Db.Sslmode)

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatalln("Не удается спарсить конфигурацию:", err)
	}
	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(10)
	return PostPostgressRepository{db}
}
