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

	postDelivery.Time = new(time.Time)
	*postDelivery.Time = time.Now()

	postDelivery.Id = new(uuid.UUID)
	*postDelivery.Id = uuid.New()

	code, err := h.postRepo.Create(ctx, postDelivery)
	return *postDelivery.Id, code, err
}
