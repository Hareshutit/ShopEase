package delivery

import (
	app "github.com/Hareshutit/ShopEase/internal/user/usecase"
	"github.com/rs/zerolog"
)

func CreateGrpcServer(command app.Commands, query app.Queries, log zerolog.Logger) GrpcServer {
	return GrpcServer{
		command: command,
		query:   query,
	}
}
