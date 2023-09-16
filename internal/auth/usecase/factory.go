package usecase

import (
	"context"

	config "github.com/Hareshutit/ShopEase/config/auth"
	"github.com/Hareshutit/ShopEase/internal/auth/domain"
	"github.com/Hareshutit/ShopEase/internal/auth/usecase/commands"
)

func NewUsecase(ctx context.Context, repo domain.CUDRepository, cfg config.Config) (Commands, Queries) {
	return Commands{
			CreateAccessToken:  commands.CreateAccessTokenHandle{PrivateKey: cfg.KeyValue.Access},
			CreateRefreshToken: commands.CreateRefreshTokenHandle{PrivateKey: cfg.KeyValue.Refresh, CUDRepository: repo},
			UpdateRefreshToken: commands.UpdateRefreshTokenHandle{PrivateKey: cfg.KeyValue.Refresh, CUDRepository: repo},
			DeleteRefreshToken: commands.DeleteRefreshTokenHandle{PrivateKey: cfg.KeyValue.Refresh, CUDRepository: repo},
		},
		Queries{}
}
