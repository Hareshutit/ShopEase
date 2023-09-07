package command

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DeleteHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *DeleteHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (int, error) {
	return h.postRepo.Delete(ctx, id)
}

type RemoveCartHandler struct {
	cartRepo  domain.CUDCartRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *RemoveCartHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
	postId uuid.UUID,
) (int, error) {
	return h.cartRepo.Remove(ctx, userId, postId)
}

type RemoveFavoriteHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *RemoveFavoriteHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
	postId uuid.UUID,
) (int, error) {
	return h.postRepo.RemoveFavorite(ctx, userId, postId)
}
