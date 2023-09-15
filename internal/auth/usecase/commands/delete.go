package commands

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"net/http"

	"github.com/Hareshutit/ShopEase/internal/auth/domain"
	"github.com/lestrrat-go/jwx/jwt"
)

type DeleteRefreshTokenHandle struct {
	PrivateKey    *ecdsa.PrivateKey
	CUDRepository domain.CUDRepository
}

func (f *DeleteRefreshTokenHandle) Delete(ctx context.Context, token string) (int, error) {
	cloneToken, err := jwt.ParseString(token)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	idUser, ok := cloneToken.Get("sub")
	if ok != true {
		return http.StatusBadRequest, fmt.Errorf("Not finding id")
	}
	idToken, ok := cloneToken.Get("jti")
	if ok != true {
		return http.StatusBadRequest, fmt.Errorf("Not finding idToken")
	}

	sidUser, ok := idUser.(string)
	if ok != true {
		return http.StatusBadRequest, fmt.Errorf("Bad id")
	}
	sidToken, ok := idToken.(string)
	if ok != true {
		return http.StatusBadRequest, fmt.Errorf("Bad idToken")
	}

	f.CUDRepository.Delete(ctx, sidUser, sidToken)
	return http.StatusOK, err
}
