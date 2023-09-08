package query

import (
	"github.com/Hareshutit/ShopEase/internal/post/domain"

	"github.com/sirupsen/logrus"
)

func NewGetMiniPostSortNewHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetMiniPostSortNewHandler {
	return GetMiniPostSortNewHandler{postRepo: postRepo, loger: loger}
}

func NewGetIdHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetIdHandler {
	return GetIdHandler{postRepo: postRepo, loger: loger}
}
