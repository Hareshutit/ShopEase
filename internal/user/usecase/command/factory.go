package command

import (
	"github.com/Hareshutit/ShopEase/internal/user/domain"
	"github.com/rs/zerolog"
)

func NewCreateUserHandler(
	userRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *zerolog.Logger,
) CreateUserHandler {
	return CreateUserHandler{userRepo: userRepo,
		validator: validator, loger: loger}
}

func NewUpdateUserHandler(
	userRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *zerolog.Logger,
) UpdateUserHandler {
	return UpdateUserHandler{userRepo: userRepo,
		validator: validator, loger: loger}
}

func NewDeleteUserHandler(
	userRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *zerolog.Logger,
) DeleteUserHandler {
	return DeleteUserHandler{userRepo: userRepo, validator: validator, loger: loger}
}
