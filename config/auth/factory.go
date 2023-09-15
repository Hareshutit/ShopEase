package config

func CreateConfig() Config {
	return Config{
		Http:  CreateHttpConfig(),
		Grcp:  CreateGrcpConfig(),
		Redis: CreateRedisConfig(),
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
		Host: "localhost",
		Port: 5432,
	}
}
