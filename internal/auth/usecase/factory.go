package usecase

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/auth/domain"
	"github.com/Hareshutit/ShopEase/internal/auth/usecase/commands"

	"github.com/deepmap/oapi-codegen/pkg/ecdsafile"
)

const PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIN2dALnjdcZaIZg4QuA6Dw+kxiSW502kJfmBN3priIhPoAoGCCqGSM49
AwEHoUQDQgAE4pPyvrB9ghqkT1Llk0A42lixkugFd/TBdOp6wf69O9Nndnp4+HcR
s9SlG/8hjB2Hz42v4p3haKWv3uS1C6ahCQ==
-----END EC PRIVATE KEY-----`

func NewUsecase(ctx context.Context, repo domain.CUDRepository) (Commands, Queries) {
	privKey, _ := ecdsafile.LoadEcdsaPrivateKey([]byte(PrivateKey))

	return Commands{
			CreateAccessToken:  commands.CreateAccessTokenHandle{PrivateKey: privKey},
			CreateRefreshToken: commands.CreateRefreshTokenHandle{PrivateKey: privKey, CUDRepository: repo},
			UpdateRefreshToken: commands.UpdateRefreshTokenHandle{PrivateKey: privKey, CUDRepository: repo},
			DeleteRefreshToken: commands.DeleteRefreshTokenHandle{PrivateKey: privKey, CUDRepository: repo},
		},
		Queries{}
}
