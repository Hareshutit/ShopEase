package command

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/user/domain"
	"github.com/rs/zerolog"

	"github.com/google/uuid"
)

type DeleteUserHandler struct {
	userRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *zerolog.Logger
}

func (h *DeleteUserHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := h.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
