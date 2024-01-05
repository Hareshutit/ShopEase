package query

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/user/domain"
	"github.com/rs/zerolog"

	"github.com/google/uuid"
)

type GetUserHandler struct {
	userRepo domain.RRepository
	loger    *zerolog.Logger
}

func (h GetUserHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (domain.User, error) {
	return h.userRepo.Get(ctx, id)
}

type FindByIdUserHandler struct {
	userRepo domain.RRepository
	loger    *zerolog.Logger
}

func (h FindByIdUserHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (domain.User, error) {
	return h.userRepo.FindById(ctx, id)
}
