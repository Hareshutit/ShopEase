package config

import "crypto/ecdsa"

type Config struct {
	Http     HttpConfig
	Grcp     GrcpConfig
	Redis    RedisConfig
	KeyValue KeyValue
}

type HttpConfig struct {
	Port int
}

type GrcpConfig struct {
	Port int
}

type RedisConfig struct {
	Host string
	Port int
}

type KeyValue struct {
	Refresh *ecdsa.PrivateKey
	Access  *ecdsa.PrivateKey
}
