package v2

import app "github.com/Hareshutit/ShopEase/internal/auth/usecase"

func CreateHttpServer(command app.Commands, query app.Queries) HttpServer {
	return HttpServer{
		command: command,
		query:   query,
	}
}
