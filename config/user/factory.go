package config

import (
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/ecdsafile"
)

const PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIN2dALnjdcZaIZg4QuA6Dw+kxiSW502kJfmBN3priIhPoAoGCCqGSM49
AwEHoUQDQgAE4pPyvrB9ghqkT1Llk0A42lixkugFd/TBdOp6wf69O9Nndnp4+HcR
s9SlG/8hjB2Hz42v4p3haKWv3uS1C6ahCQ==
-----END EC PRIVATE KEY-----`

func CreateConfig() Config {
	return Config{
		Http:          CreateHttpConfig(),
		Grcp:          CreateGrcpConfig(),
		Db:            CreateDataBaseConfig(),
		Authorization: CreateSecurityConfig(),
	}
}

func CreateHttpConfig() HttpConfig {
	return HttpConfig{8080}
}

func CreateGrcpConfig() GrcpConfig {
	return GrcpConfig{8081}
}

func CreateDataBaseConfig() DataBaseConfig {
	return DataBaseConfig{
		User:         "shopease",
		DataBaseName: "user",
		Password:     "uniq123",
		Host:         "postgres_user",
		Port:         5432,
		Sslmode:      "disable",
	}
}

func CreateSecurityConfig() SecurityConfig {
	aprivatekey, err := ecdsafile.LoadEcdsaPrivateKey([]byte(PrivateKey))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return SecurityConfig{
		Verify: &aprivatekey.PublicKey,
	}
}
