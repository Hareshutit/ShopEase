package command

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/post/domain"
	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

type UpdateHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *UpdateHandler) Handle(
	ctx context.Context,
	postDelivery domain.Post,
) (int, error) {
	return h.postRepo.Update(ctx, postDelivery)
}

type IncrementViewsHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *IncrementViewsHandler) Handle(
	ctx context.Context,
	PostId uuid.UUID,
	UserId uuid.UUID,
) (int, error) {
	return h.postRepo.IncrementViews(ctx, PostId, UserId)
}
