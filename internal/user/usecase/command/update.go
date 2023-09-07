package command

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/user/domain"

	"github.com/sirupsen/logrus"
)

type UpdateUserHandler struct {
	userRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *UpdateUserHandler) Handle(
	ctx context.Context,
	userDelivery domain.User,
) error {

	user := domain.User{
		Id: userDelivery.Id,

		Email:       userDelivery.Email,
		PhoneNumber: userDelivery.PhoneNumber,

		Login:    userDelivery.Login,
		Password: userDelivery.Password,

		Name: userDelivery.Name,

		PathToAvatar: userDelivery.PathToAvatar,
	}
	err := h.userRepo.Update(ctx, user)
	return err
}
