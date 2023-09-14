package commands

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

func createToken(claims map[string]any, PrivateKey *ecdsa.PrivateKey) ([]byte, int, error) {
	t := jwt.New()
	for tag, claim := range claims {
		err := t.Set(tag, claim)
		if err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("setting payload: %w", err)
		}
	}
	return sign(t, PrivateKey)
}

func sign(t jwt.Token, PrivateKey *ecdsa.PrivateKey) ([]byte, int, error) {
	hdr := jws.NewHeaders()
	if err := hdr.Set(jws.AlgorithmKey, jwa.ES256); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("setting algorithm: %w", err)
	}
	if err := hdr.Set(jws.TypeKey, "JWT"); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("setting type: %w", err)
	}
	token, err := jwt.Sign(t, jwa.ES256, PrivateKey, jwt.WithHeaders(hdr))
	return token, http.StatusOK, err
}
