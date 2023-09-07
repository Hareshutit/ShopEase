package commands

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

type CreateJWSHandle struct {
	PrivateKey *ecdsa.PrivateKey
	KeySet     jwk.Set
}

func (f *CreateJWSHandle) CreateJWSWithClaims(claims map[string]string, audience string, issuer string) ([]byte, error) {
	t := jwt.New()
	for tag, claim := range claims {
		err := t.Set(tag, claim)
		if err != nil {
			return nil, fmt.Errorf("setting payload: %w", err)
		}
	}
	return f.SignToken(t)
}

func (f *CreateJWSHandle) SignToken(t jwt.Token) ([]byte, error) {
	hdr := jws.NewHeaders()
	if err := hdr.Set(jws.AlgorithmKey, jwa.ES256); err != nil {
		return nil, fmt.Errorf("setting algorithm: %w", err)
	}
	if err := hdr.Set(jws.TypeKey, "JWT"); err != nil {
		return nil, fmt.Errorf("setting type: %w", err)
	}
	if err := hdr.Set(jws.KeyIDKey, `key-id`); err != nil {
		return nil, fmt.Errorf("setting Key ID: %w", err)
	}
	return jwt.Sign(t, jwa.ES256, f.PrivateKey, jwt.WithHeaders(hdr))
}
