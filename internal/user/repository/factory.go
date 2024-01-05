package repository

import (
	"database/sql"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/rs/zerolog"
)

func CreatePostgressRepository(dsn string, log zerolog.Logger) UserPostgressRepository {
	db, err := sql.Open("pgx", dsn)
	log = log.With().Str("Layer", "Repository").Logger()
	if err != nil {
		log.Fatal().Msg("Failed to retrieve the configuration\n: " + err.Error())
	}
	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	db.SetMaxOpenConns(10)
	return UserPostgressRepository{db, log}
}
