package config

import "crypto/ecdsa"

type Config struct {
	Http          HttpConfig
	Grcp          GrcpConfig
	Db            DataBaseConfig
	Authorization SecurityConfig
}

type HttpConfig struct {
	Port int
}

type GrcpConfig struct {
	Port int
}

type DataBaseConfig struct {
	User         string
	DataBaseName string
	Password     string
	Host         string
	Port         int
	Sslmode      string
}

type SecurityConfig struct {
	Verify *ecdsa.PublicKey
}
