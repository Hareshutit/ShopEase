package v2

import (
	app "github.com/Hareshutit/ShopEase/internal/user/usecase"
	"github.com/rs/zerolog"
)

func CreateHttpServer(command app.Commands, query app.Queries, log zerolog.Logger) HttpServer {
	log = log.With().Str("Layer", "Delivery").Str("Protocol", "Http").Logger()
	return HttpServer{
		command: command,
		query:   query,
		log:     log,
	}
}
