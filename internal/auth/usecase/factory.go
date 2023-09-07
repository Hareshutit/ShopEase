package usecase

import (
	"context"

	"github.com/Hareshutit/ShopEase/internal/auth/usecase/commands"

	"github.com/deepmap/oapi-codegen/pkg/ecdsafile"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
)

const KeyID = `key-id`

const PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIN2dALnjdcZaIZg4QuA6Dw+kxiSW502kJfmBN3priIhPoAoGCCqGSM49
AwEHoUQDQgAE4pPyvrB9ghqkT1Llk0A42lixkugFd/TBdOp6wf69O9Nndnp4+HcR
s9SlG/8hjB2Hz42v4p3haKWv3uS1C6ahCQ==
-----END EC PRIVATE KEY-----`

func NewUsecase(ctx context.Context) (Commands, Queries) {
	privKey, _ := ecdsafile.LoadEcdsaPrivateKey([]byte(PrivateKey))

	set := jwk.NewSet()
	pubKey := jwk.NewECDSAPublicKey()

	_ = pubKey.FromRaw(&privKey.PublicKey)

	_ = pubKey.Set(jwk.AlgorithmKey, jwa.ES256)

	_ = pubKey.Set(jwk.KeyIDKey, KeyID)

	set.Add(pubKey)
	return Commands{
			CreateToken: commands.CreateJWSHandle{KeySet: set, PrivateKey: privKey},
		},
		Queries{}
}
