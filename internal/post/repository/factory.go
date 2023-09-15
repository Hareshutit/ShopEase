package repository

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/Hareshutit/ShopEase/config/post"
	_ "github.com/jackc/pgx/stdlib"
)

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
