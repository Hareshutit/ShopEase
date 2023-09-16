package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
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
		Http:     CreateHttpConfig(),
		Grcp:     CreateGrcpConfig(),
		Redis:    CreateRedisConfig(),
		KeyValue: CreateKey(),
	}
}

func CreateHttpConfig() HttpConfig {
	return HttpConfig{8083}
}

func CreateGrcpConfig() GrcpConfig {
	return GrcpConfig{8084}
}

func CreateRedisConfig() RedisConfig {
	return RedisConfig{
		Host: "redis_auth",
		Port: 6379,
	}
}

func CreateKey() KeyValue {
	pubkeyCurve := elliptic.P256()
	rprivatekey := new(ecdsa.PrivateKey)
	rprivatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	aprivatekey, _ := ecdsafile.LoadEcdsaPrivateKey([]byte(PrivateKey))
	return KeyValue{
		Refresh: rprivatekey,
		Access:  aprivatekey,
	}
}
