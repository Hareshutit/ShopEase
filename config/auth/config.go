package config

type Config struct {
	Http  HttpConfig
	Grcp  GrcpConfig
	Redis RedisConfig
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
