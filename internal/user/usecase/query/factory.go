package query

import (
	"github.com/Hareshutit/ShopEase/internal/user/domain"

	"github.com/sirupsen/logrus"
)

func NewGetUserHandler(
	userRepo domain.RRepository,
	loger *logrus.Entry,
) GetUserHandler {
	return GetUserHandler{userRepo: userRepo, loger: loger}
}

func NewFindByIdUserHandler(
	userRepo domain.RRepository,
	loger *logrus.Entry,
) FindByIdUserHandler {
	return FindByIdUserHandler{userRepo: userRepo, loger: loger}
}

func NewCheckUserHandler(
	userRepo domain.RRepository,
	loger *logrus.Entry,
) CheckUserHandler {
	return CheckUserHandler{userRepo: userRepo, loger: loger}
}
