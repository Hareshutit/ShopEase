package usecase

import (
	"github.com/Hareshutit/ShopEase/internal/post/usecase/command"
	"github.com/Hareshutit/ShopEase/internal/post/usecase/query"
)

// В данном агрегате перечисленны все команды сервиса объявлений
type Commands struct {
	CreatePost command.CreateHandler
	UpdatePost command.UpdateHandler
	DeletePost command.DeleteHandler
}

// В данном агрегате перечисленны все запросы сервиса объявлений
type Queries struct {
	GetIdPost          query.GetIdHandler
	GetMiniPostSortNew query.GetMiniPostSortNewHandler
}
