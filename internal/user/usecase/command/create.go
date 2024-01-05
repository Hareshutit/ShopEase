package command

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/user/domain"
	"github.com/rs/zerolog"

	"github.com/google/uuid"
)

type CreateUserHandler struct {
	userRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *zerolog.Logger
}

func (h *CreateUserHandler) Handle(
	ctx context.Context,
	passwordCheck string,
	userDelivery domain.User,
) (*uuid.UUID, error) {
	if userDelivery.Password != passwordCheck {
		return nil, domain.PassNonComporableErr{}
	}

	if err := h.validator.Email.IsValid(userDelivery.Email); err != nil {
		return nil, err
	}
	if err := h.validator.Login.IsValid(userDelivery.Login); err != nil {
		return nil, err
	}
	if err := h.validator.PhoneNumber.IsValid(userDelivery.PhoneNumber); err != nil {
		return nil, err
	}
	if err := h.validator.Password.IsValid(userDelivery.Password); err != nil {
		return nil, err
	}
	if err := h.validator.Name.IsValid(userDelivery.Name); err != nil {
		return nil, err
	}
	if err := h.validator.Avatar.IsValid(userDelivery.PathToAvatar); err != nil {
		return nil, err
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
	if err != nil {
		return nil, err
	}

	return &user.Id, nil
}
