package usecase

import (
	"context"

	config "github.com/Hareshutit/ShopEase/config/post"
	"github.com/Hareshutit/ShopEase/internal/post/domain"
	"github.com/Hareshutit/ShopEase/internal/post/repository"
	"github.com/Hareshutit/ShopEase/internal/post/usecase/command"
	"github.com/Hareshutit/ShopEase/internal/post/usecase/query"

	"github.com/sirupsen/logrus"
)

func NewUsecase(ctx context.Context, cfg config.Config) (Commands, Queries) {
	validator := domain.CreateSpecificationManager(cfg)

	postRepository := repository.CreatePostgressRepository(cfg)

	logger := logrus.NewEntry(logrus.StandardLogger())

	return Commands{
			Create: command.NewCreateHandler(&postRepository, validator, logger),
			Update: command.NewUpdateHandler(&postRepository, validator, logger),
			Delete: command.NewDeleteHandler(&postRepository, validator, logger),
		},
		Queries{
			GetById:       query.NewGetByIdHandler(postRepository, logger),
			GetMiniObject: query.NewGetMiniObjectHandler(postRepository, logger),
		}
}
