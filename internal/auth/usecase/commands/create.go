package commands

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/Hareshutit/ShopEase/internal/auth/domain"
	"github.com/google/uuid"
)

type CreateRefreshTokenHandle struct {
	PrivateKey    *ecdsa.PrivateKey
	CUDRepository domain.CUDRepository
}

func (f *CreateRefreshTokenHandle) Create(ctx context.Context, id string) ([]byte, int, error) {
	claims := make(map[string]any)
	claims["iss"] = "auth.shopease.com"
	claims["aud"] = "shopease.com"
	claims["sub"] = id
	idToken := uuid.New().String()
	claims["jti"] = idToken
	claims["exp"] = time.Now().Unix() + 2592000

	token, code, err := createToken(claims, f.PrivateKey)
	code, err = f.CUDRepository.Create(ctx, id, idToken, token)
	if err != nil {
		return nil, code, err
	}
	return token, code, err
}

type CreateAccessTokenHandle struct {
	PrivateKey *ecdsa.PrivateKey
}

func (f *CreateAccessTokenHandle) Create(id string) ([]byte, int, error) {
	claims := make(map[string]any)
	claims["iss"] = "auth.shopease.com"
	claims["aud"] = "shopease.com"
	claims["sub"] = id
	claims["exp"] = time.Now().Unix() + 3600

	return createToken(claims, f.PrivateKey)
}
