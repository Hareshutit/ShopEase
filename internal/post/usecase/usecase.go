package usecase

import (
	"github.com/Hareshutit/ShopEase/internal/post/usecase/command"
	"github.com/Hareshutit/ShopEase/internal/post/usecase/query"
)

// В данном агрегате перечисленны все команды сервиса объявлений
type Commands struct {
	Create command.CreateHandler
	Update command.UpdateHandler
	Delete command.DeleteHandler
}

// В данном агрегате перечисленны все запросы сервиса объявлений
type Queries struct {
	GetById       query.GetByIdHandler
	GetMiniObject query.GetMiniObjectHandler
}
