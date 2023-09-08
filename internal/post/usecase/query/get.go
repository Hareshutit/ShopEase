package query

import (
	"context"
	"net/http"

	"github.com/Hareshutit/ShopEase/internal/post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetMiniPostSortNewHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetMiniPostSortNewHandler) Handle(
	ctx context.Context,
	postParam domain.Parameters,
) ([]domain.Post, int, error) {
	return h.postRepo.GetMiniPostSortNew(ctx, postParam)
}

type GetIdHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetIdHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Post, int, error) {
	return h.postRepo.GetIdPost(ctx, id)
}

type GetCartHandler struct {
	cartRepo domain.RCartRepository
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetCartHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
) ([]domain.Post, int, error) {
	uuids, code, err := h.cartRepo.GetUUID(ctx, userId)
	if code != http.StatusOK {
		return nil, code, err
	}
	return h.postRepo.GetCart(ctx, uuids)
}

type GetFavoriteHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetFavoriteHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
) ([]uuid.UUID, int, error) {
	return h.postRepo.GetFavorite(ctx, userId)
}

type SearchPostHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h SearchPostHandler) Handle(
	ctx context.Context,
	search string,
) ([]uuid.UUID, int, error) {
	return h.postRepo.SearchPost(ctx, search)
}
