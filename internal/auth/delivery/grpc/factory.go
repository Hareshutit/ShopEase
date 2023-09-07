package delivery

import app "github.com/Hareshutit/ShopEase/internal/auth/usecase"

func CreateGrpcServer(command app.Commands, query app.Queries) GrpcServer {
	return GrpcServer{
		command: command,
		query:   query,
	}
}
