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

func CreateValidConfig() ValidConfig {
	return ValidConfig{
		LoginValidate:      CreateLoginConfig(),
		PasswordValidate:   CreatePasswordConfig(),
		SecondNameValidate: CreateSecondNameConfig(),
		FirstNameValidate:  CreateFirstNameConfig(),
		PatronimicValidate: CreatePatronimicConfig(),
		AvatarValidate:     CreateAvatarConfig(),
	}
}

func CreatePasswordConfig() PasswordConfig {
	SpecialChar := make(map[rune]bool)
	SpecialChar['/'] = true
	SpecialChar['.'] = true
	SpecialChar[','] = true
	return PasswordConfig{SpecialChar, 0, 20}
}

func CreateLoginConfig() LoginConfig {
	NonAcceptableValues := make(map[rune]bool)
	NonAcceptableValues['/'] = true
	NonAcceptableValues['.'] = true
	NonAcceptableValues[','] = true
	NonAcceptableValues['@'] = true
	return LoginConfig{0, 20, NonAcceptableValues}
}

func CreateSecondNameConfig() SecondNameConfig {
	return SecondNameConfig{0, 20}
}

func CreateFirstNameConfig() FirstNameConfig {
	return FirstNameConfig{0, 20}
}

func CreatePatronimicConfig() PatronimicConfig {
	return PatronimicConfig{0, 20}
}

func CreateAvatarConfig() AvatarConfig {
	return AvatarConfig{20971520}
}
