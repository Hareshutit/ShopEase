package usecase

import (
	"github.com/Hareshutit/ShopEase/internal/user/usecase/command"
	"github.com/Hareshutit/ShopEase/internal/user/usecase/query"
)

// В данном агрегате перечисленны все команды сервиса пользователя
type Commands struct {
	CreateUser command.CreateUserHandler
	UpdateUser command.UpdateUserHandler
	DeleteUser command.DeleteUserHandler
}

// В данном агрегате перечисленны все запросы сервиса пользователя
type Queries struct {
	GetUser      query.GetUserHandler
	CheckUser    query.CheckUserHandler
	FindByIdUser query.FindByIdUserHandler
}
