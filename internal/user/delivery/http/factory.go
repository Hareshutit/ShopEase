package v2

import app "github.com/Hareshutit/ShopEase/internal/user/usecase"

func CreateHttpServer(command app.Commands, query app.Queries) HttpServer {
	return HttpServer{
		command: command,
		query:   query,
	}
}
