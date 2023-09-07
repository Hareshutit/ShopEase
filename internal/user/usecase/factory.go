package usecase

import (
	"context"
	"fmt"

	"github.com/Hareshutit/ShopEase/internal/user/domain"
	"github.com/Hareshutit/ShopEase/internal/user/repository"
	"github.com/Hareshutit/ShopEase/internal/user/usecase/command"
	"github.com/Hareshutit/ShopEase/internal/user/usecase/query"

	config "github.com/Hareshutit/ShopEase/config/user"

	"github.com/sirupsen/logrus"
)

func NewUsecase(ctx context.Context, cfg config.Config) (Commands, Queries) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Db.User, cfg.Db.DataBaseName, cfg.Db.Password, cfg.Db.Host,
		cfg.Db.Port, cfg.Db.Sslmode)
	validator := domain.CreateSpecificationManager(cfg)
	Repository := repository.CreatePostgressRepository(dsn)
	logger := logrus.NewEntry(logrus.StandardLogger())

	return Commands{
			CreateUser: command.NewCreateUserHandler(&Repository, validator, logger),
			UpdateUser: command.NewUpdateUserHandler(&Repository, validator, logger),
			DeleteUser: command.NewDeleteUserHandler(&Repository, validator, logger),
		},
		Queries{
			GetUser:      query.NewGetUserHandler(Repository, logger),
			CheckUser:    query.NewCheckUserHandler(Repository, logger),
			FindByIdUser: query.NewFindByIdUserHandler(Repository, logger),
		}
}
