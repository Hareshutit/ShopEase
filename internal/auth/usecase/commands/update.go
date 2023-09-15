package commands

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"net/http"

	"github.com/Hareshutit/ShopEase/internal/auth/domain"
	"github.com/lestrrat-go/jwx/jwt"
)

type UpdateRefreshTokenHandle struct {
	PrivateKey    *ecdsa.PrivateKey
	CUDRepository domain.CUDRepository
}

func (f *UpdateRefreshTokenHandle) Update(ctx context.Context, token string) ([]byte, *string, int, error) {
	cloneToken, err := jwt.ParseString(token)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	idUser, ok := cloneToken.Get("sub")
	if ok != true {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("Not finding id")
	}
	idToken, ok := cloneToken.Get("jti")
	if ok != true {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("Not finding idToken")
	}

	sidUser, ok := idUser.(string)
	if ok != true {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("Bad id")
	}
	sidToken, ok := idToken.(string)
	if ok != true {
		return nil, nil, http.StatusBadRequest, fmt.Errorf("Bad idToken")
	}

	scloneToketn, code, err := sign(cloneToken, f.PrivateKey)

	f.CUDRepository.Update(ctx, sidUser, sidToken, scloneToketn)
	return scloneToketn, &sidUser, code, err
}
