package usecase

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/user/domain"
	"github.com/Hareshutit/ShopEase/internal/user/usecase/command"
	"github.com/Hareshutit/ShopEase/internal/user/usecase/query"
	"github.com/rs/zerolog"
)

func NewUsecase(ctx context.Context,
	RRepository domain.RRepository,
	CUDRepository domain.CUDRepository,
	log zerolog.Logger) (Commands, Queries) {

	validator := domain.CreateSpecificationManager()
	log = log.With().Str("Layer", "Usecase").Logger()

	return Commands{
			CreateUser: command.NewCreateUserHandler(CUDRepository, validator, &log),
			UpdateUser: command.NewUpdateUserHandler(CUDRepository, validator, &log),
			DeleteUser: command.NewDeleteUserHandler(CUDRepository, validator, &log),
		},
		Queries{
			GetUser:      query.NewGetUserHandler(RRepository, &log),
			CheckUser:    query.NewCheckUserHandler(RRepository, &log),
			FindByIdUser: query.NewFindByIdUserHandler(RRepository, &log),
		}
}
