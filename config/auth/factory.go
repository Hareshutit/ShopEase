package config

func CreateConfig() Config {
	return Config{
		Http:  CreateHttpConfig(),
		Grcp:  CreateGrcpConfig(),
		Db:    CreateDataBaseConfig(),
		Valid: CreateValidConfig(),
	}
}

func CreateHttpConfig() HttpConfig {
	return HttpConfig{8083}
}

func CreateGrcpConfig() GrcpConfig {
	return GrcpConfig{8084}
}

func CreateDataBaseConfig() DataBaseConfig {
	return DataBaseConfig{
		User:         "shopease",
		DataBaseName: "post",
		Password:     "uniq123",
		Host:         "localhost",
		Port:         5432,
		Sslmode:      "disable",
	}
}

func CreateValidConfig() ValidConfig {
	return ValidConfig{
		TitleValidate:       CreateTitleConfig(),
		DescriptionValidate: CreateDescriptionConfig(),
		ImagesValidate:      CreateImagesConfig(),
	}
}
func CreateTitleConfig() TitleConfig {
	return TitleConfig{5, 40}
}

func CreateDescriptionConfig() DescriptionConfig {
	return DescriptionConfig{20, 150}
}

func CreateImagesConfig() ImagesConfig {
	return ImagesConfig{1, 5, 20971520}
}

func CreateTagsConfig() TagsConfig {
	return TagsConfig{1, 5}
}
