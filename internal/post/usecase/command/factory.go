package command

import (
	"github.com/Hareshutit/ShopEase/internal/post/domain"

	"github.com/sirupsen/logrus"
)

func NewCreateHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) CreateHandler {
	return CreateHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewUpdateHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) UpdateHandler {
	return UpdateHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewDeleteHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) DeleteHandler {
	return DeleteHandler{postRepo: postRepo, validator: validator, loger: loger}
}

func NewAddCartHandler(
	cartRepo domain.CUDCartRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) AddCartHandler {
	return AddCartHandler{cartRepo: cartRepo, validator: validator, loger: loger}
}

func NewRemoveCartHandler(
	cartRepo domain.CUDCartRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) RemoveCartHandler {
	return RemoveCartHandler{cartRepo: cartRepo, validator: validator, loger: loger}
}

func NewAddFavoriteHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) AddFavoriteHandler {
	return AddFavoriteHandler{postRepo: postRepo, validator: validator, loger: loger}
}

func NewRemoveFavoriteHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) RemoveFavoriteHandler {
	return RemoveFavoriteHandler{postRepo: postRepo, validator: validator, loger: loger}
}
