package query

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetMiniObjectHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetMiniObjectHandler) Handle(
	ctx context.Context,
	postParam domain.Parameters,
) ([]domain.Post, int, error) {
	return h.postRepo.GetMiniObject(ctx, postParam)
}

type GetByIdHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetByIdHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Post, int, error) {
	return h.postRepo.GetById(ctx, id)
}
