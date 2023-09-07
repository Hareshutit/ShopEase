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

func NewGetCartHandler(
	cartRepo domain.RCartRepository,
	postRepository domain.RRepository,
	loger *logrus.Entry,
) GetCartHandler {
	return GetCartHandler{cartRepo: cartRepo, loger: loger}
}

func NewGetFavoriteHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetFavoriteHandler {
	return GetFavoriteHandler{postRepo: postRepo, loger: loger}
}

func NewSearchPostHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) SearchPostHandler {
	return SearchPostHandler{postRepo: postRepo, loger: loger}
}
