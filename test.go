// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/ecdsafile"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

type createJWSHandle struct {
	PrivateKey *ecdsa.PrivateKey
}

func (f *createJWSHandle) createJWSWithClaims(claims map[string]string, audience string, issuer string) ([]byte, error) {
	t := jwt.New()
	for tag, claim := range claims {
		err := t.Set(tag, claim)
		if err != nil {
			return nil, fmt.Errorf("setting payload: %w", err)
		}
	}
	return f.signToken(t)
}

func (f *createJWSHandle) signToken(t jwt.Token) ([]byte, error) {
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

const KeyID = `key-id`

const PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIN2dALnjdcZaIZg4QuA6Dw+kxiSW502kJfmBN3priIhPoAoGCCqGSM49
AwEHoUQDQgAE4pPyvrB9ghqkT1Llk0A42lixkugFd/TBdOp6wf69O9Nndnp4+HcR
s9SlG/8hjB2Hz42v4p3haKWv3uS1C6ahCQ==
-----END EC PRIVATE KEY-----`

func main() {
	claims := make(map[string]string)
	claims["id"] = "132"
	privKey, _ := ecdsafile.LoadEcdsaPrivateKey([]byte(PrivateKey))
	wd := createJWSHandle{PrivateKey: privKey}
	result, _ := wd.createJWSWithClaims(claims, "shopEaseFront", "auth")
	fmt.Println(string(result))
	clone, _ := jwt.Parse(result)
	cloneb, _ := wd.signToken(clone)
	fmt.Println(string(cloneb))
}
