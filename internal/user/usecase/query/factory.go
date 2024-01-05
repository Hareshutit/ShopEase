package query

import (
	"github.com/Hareshutit/ShopEase/internal/user/domain"
	"github.com/rs/zerolog"
)

func NewGetUserHandler(
	userRepo domain.RRepository,
	loger *zerolog.Logger,
) GetUserHandler {
	return GetUserHandler{userRepo: userRepo, loger: loger}
}

func NewFindByIdUserHandler(
	userRepo domain.RRepository,
	loger *zerolog.Logger,
) FindByIdUserHandler {
	return FindByIdUserHandler{userRepo: userRepo, loger: loger}
}

func NewCheckUserHandler(
	userRepo domain.RRepository,
	loger *zerolog.Logger,
) CheckUserHandler {
	return CheckUserHandler{userRepo: userRepo, loger: loger}
}
