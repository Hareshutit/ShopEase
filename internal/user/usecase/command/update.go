package command

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/user/domain"
	"github.com/rs/zerolog"
)

type UpdateUserHandler struct {
	userRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *zerolog.Logger
}

func (h *UpdateUserHandler) Handle(
	ctx context.Context,
	userDelivery domain.User,
) error {

	if err := h.validator.Email.IsValid(userDelivery.Email); err != nil {
		return err
	}
	if err := h.validator.Login.IsValid(userDelivery.Login); err != nil {
		return err
	}
	if err := h.validator.PhoneNumber.IsValid(userDelivery.PhoneNumber); err != nil {
		return err
	}
	if err := h.validator.Password.IsValid(userDelivery.Password); err != nil {
		return err
	}
	if err := h.validator.Name.IsValid(userDelivery.Name); err != nil {
		return err
	}
	if err := h.validator.Avatar.IsValid(userDelivery.PathToAvatar); err != nil {
		return err
	}

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
	if err != nil {
		return err
	}

	return nil
}
