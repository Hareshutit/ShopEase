package command

import (
	"context"
	"time"

	"github.com/Hareshutit/ShopEase/internal/post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *CreateHandler) Handle(
	ctx context.Context,
	postDelivery domain.Post,
) (uuid.UUID, int, error) {
	*postDelivery.Time = time.Now()
	*postDelivery.Id = uuid.New()
	code, err := h.postRepo.Create(ctx, postDelivery)
	return *postDelivery.Id, code, err
}

type AddCartHandler struct {
	cartRepo  domain.CUDCartRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *AddCartHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
	postId uuid.UUID,
) (int, error) {
	code, err := h.cartRepo.Add(ctx, userId, postId)
	return code, err
}

type AddFavoriteHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *AddFavoriteHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
	postId uuid.UUID,
) (int, error) {
	code, err := h.postRepo.AddFavorite(ctx, userId, postId)
	return code, err
}
