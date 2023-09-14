package usecase

import "github.com/Hareshutit/ShopEase/internal/auth/usecase/commands"

// В данном агрегате перечисленны все команды сервиса авторизации
type Commands struct {
	CreateAccessToken  commands.CreateAccessTokenHandle
	CreateRefreshToken commands.CreateRefreshTokenHandle
	UpdateRefreshToken commands.UpdateRefreshTokenHandle
}

// В данном агрегате перечисленны все запросы сервиса авторизации
type Queries struct {
}
