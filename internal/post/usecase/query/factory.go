package query

import (
	"github.com/Hareshutit/ShopEase/internal/post/domain"

	"github.com/sirupsen/logrus"
)

func NewGetMiniObjectHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetMiniObjectHandler {
	return GetMiniObjectHandler{postRepo: postRepo, loger: loger}
}

func NewGetByIdHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetByIdHandler {
	return GetByIdHandler{postRepo: postRepo, loger: loger}
}
