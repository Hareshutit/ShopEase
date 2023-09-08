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
