package command

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/user/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DeleteUserHandler struct {
	userRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *DeleteUserHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := h.userRepo.Delete(ctx, id)
	return err
}
