package user

import (
	"crypto/ecdsa"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

const PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIN2dALnjdcZaIZg4QuA6Dw+kxiSW502kJfmBN3priIhPoAoGCCqGSM49
AwEHoUQDQgAE4pPyvrB9ghqkT1Llk0A42lixkugFd/TBdOp6wf69O9Nndnp4+HcR
s9SlG/8hjB2Hz42v4p3haKWv3uS1C6ahCQ==
-----END EC PRIVATE KEY-----`

const KeyID = `key-id`

const PermissionsClaim = "perm"

type InstanceAuthenticator struct {
	key      *ecdsa.PrivateKey
	alg      jwa.SignatureAlgorithm
	audience string
	issuer   string
}

func NewInstanceAuthenticator(key *ecdsa.PrivateKey, alg jwa.SignatureAlgorithm,
	audience string, issuer string) (*InstanceAuthenticator, error) {
	return &InstanceAuthenticator{
		key:      key,
		alg:      alg,
		audience: audience,
		issuer:   issuer,
	}, nil
}

func (f *InstanceAuthenticator) ValidateJWS(jwsString string) (jwt.Token, error) {
	return jwt.Parse([]byte(jwsString), jwt.WithVerify(f.alg, f.key),
		jwt.WithAudience(f.audience), jwt.WithIssuer(f.issuer), jwt.WithClock(jwt.ClockFunc(time.Now)))
}
