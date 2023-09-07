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

	cartRepository := repository.CreateRedisRepository(cfg)
	postRepository := repository.CreatePostgressRepository(cfg)

	logger := logrus.NewEntry(logrus.StandardLogger())

	return Commands{
			CreatePost:     command.NewCreateHandler(&postRepository, validator, logger),
			UpdatePost:     command.NewUpdateHandler(&postRepository, validator, logger),
			DeletePost:     command.NewDeleteHandler(&postRepository, validator, logger),
			AddFavorite:    command.NewAddFavoriteHandler(&postRepository, validator, logger),
			RemoveFavorite: command.NewRemoveFavoriteHandler(&postRepository, validator, logger),
			AddCart:        command.NewAddCartHandler(&cartRepository, validator, logger),
			RemoveCart:     command.NewRemoveCartHandler(&cartRepository, validator, logger),
		},
		Queries{
			GetIdPost:          query.NewGetIdHandler(postRepository, logger),
			GetMiniPostSortNew: query.NewGetMiniPostSortNewHandler(postRepository, logger),
			GetCart:            query.NewGetCartHandler(&cartRepository, postRepository, logger),
			GetFavorite:        query.NewGetFavoriteHandler(postRepository, logger),
			SearhPost:          query.NewSearchPostHandler(postRepository, logger),
		}
}
