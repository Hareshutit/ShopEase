package command

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/user/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateUserHandler struct {
	userRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *CreateUserHandler) Handle(
	ctx context.Context,
	passwordCheck string,
	userDelivery domain.User,
) (uuid.UUID, error) {
	if userDelivery.Password != passwordCheck {
		return uuid.UUID{}, domain.PassNonComporableErr{}
	}

	user := domain.User{
		Id: uuid.New(),

		Email:       userDelivery.Email,
		PhoneNumber: userDelivery.PhoneNumber,

		Login:    userDelivery.Login,
		Password: userDelivery.Password,

		Name: userDelivery.Name,

		PathToAvatar: userDelivery.PathToAvatar,
	}
	err := h.userRepo.Create(ctx, user)
	return user.Id, err
}
