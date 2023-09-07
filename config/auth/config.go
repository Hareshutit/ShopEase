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
	TitleValidate       TitleConfig
	DescriptionValidate DescriptionConfig
	ImagesValidate      ImagesConfig
	TagsValidate        TagsConfig
}

type TitleConfig struct {
	MinLength uint
	MaxLength uint
}

type DescriptionConfig struct {
	MinLength uint
	MaxLength uint
}

type ImagesConfig struct {
	MinCount uint
	MaxCount uint
	Weigth   uint
}

type TagsConfig struct {
	MinLength uint
	MaxLength uint
}
