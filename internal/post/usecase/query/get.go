package query

import (
	"context"

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
