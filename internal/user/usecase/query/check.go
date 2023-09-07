package query

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/user/domain"

	"github.com/sirupsen/logrus"
)

type CheckUserHandler struct {
	userRepo domain.RRepository
	loger    *logrus.Entry
}

func (h CheckUserHandler) Handle(
	ctx context.Context,
	login string,
	password string,
) string {
	return h.userRepo.CheckUser(ctx, login, password)
}
