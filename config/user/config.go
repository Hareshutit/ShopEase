package config

type Config struct {
	Http  HttpConfig
	Grcp  GrcpConfig
	Db    DataBaseConfig
	Valid ValidConfig
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

type ValidConfig struct {
	LoginValidate      LoginConfig
	PasswordValidate   PasswordConfig
	SecondNameValidate SecondNameConfig
	FirstNameValidate  FirstNameConfig
	PatronimicValidate PatronimicConfig
	AvatarValidate     AvatarConfig
}

type PasswordConfig struct {
	SpecialChar map[rune]bool
	MinLength   uint
	MaxLength   uint
}

type LoginConfig struct {
	MinLength           uint
	MaxLength           uint
	NonAcceptableValues map[rune]bool
}

type SecondNameConfig struct {
	MinLength uint
	MaxLength uint
}

type FirstNameConfig struct {
	MinLength uint
	MaxLength uint
}

type PatronimicConfig struct {
	MinLength uint
	MaxLength uint
}

type AvatarConfig struct {
	Weigth uint
}
