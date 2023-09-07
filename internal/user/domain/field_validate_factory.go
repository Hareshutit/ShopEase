package domain

import config "github.com/Hareshutit/ShopEase/config/user"

func CreateSpecificationManager(cfg config.Config) SpecificationManager {
	return SpecificationManager{
		Email:       CreateEmailSpecification(cfg),
		Login:       CreateLoginSpecification(cfg),
		PhoneNumber: CreatePhoneNumberSpecification(cfg),
		Password:    CreatePassSpecification(cfg),
		FirstName:   CreateFirstNameSpecification(cfg),
		SecondName:  CreateSecondNameSpecification(cfg),
		Patronimic:  CreatePatronimicSpecification(cfg),
		Avatar:      CreateAvatarSpecification(cfg),
	}
}

func CreatePassSpecification(cfg config.Config) PassSpecification {
	var valid []Specification[string]
	valid = append(valid, CreatePassAndSpecification(cfg))
	return PassSpecification{valid}
}

func CreatePassAndSpecification(cfg config.Config) PassAndSpecification {
	var valid []Specification[string]
	valid = append(valid, PassLengthValidation{cfg.Valid.PasswordValidate.MaxLength,
		cfg.Valid.PasswordValidate.MinLength})
	validFor := CreatePassForRuneValidation(cfg)
	valid = append(valid, validFor)
	return PassAndSpecification{valid}
}

func CreatePassLengthValidation(cfg config.Config) PassLengthValidation {
	return PassLengthValidation{cfg.Valid.PasswordValidate.MaxLength,
		cfg.Valid.PasswordValidate.MinLength}
}

func CreatePassForRuneValidation(cfg config.Config) PassForRuneValidation {
	validspecialCharacters := cfg.Valid.PasswordValidate.SpecialChar
	validRune := make(map[string]Specification[rune])
	validRune["SC"] = passSpecialCharactersValidation{validspecialCharacters}
	validRune["UC"] = passUpperCaseValidation{}
	validRune["LC"] = passLowerCaseValidation{}
	return PassForRuneValidation{validRune}
}

func CreateEmailSpecification(cfg config.Config) PassSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateEmailAndSpecification(cfg))
	return PassSpecification{valid}
}

func CreateEmailAndSpecification(cfg config.Config) PassAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateEmailRequiredValueValidation(cfg))
	return PassAndSpecification{valid}
}

func CreateEmailRequiredValueValidation(cfg config.Config) EmailRequiredValueValidation {
	return EmailRequiredValueValidation{}
}

func CreateLoginSpecification(cfg config.Config) LoginSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateLoginAndSpecification(cfg))
	return LoginSpecification{valid}
}

func CreateLoginAndSpecification(cfg config.Config) LoginAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateLoginAcceptableValuesValidation(cfg))
	valid = append(valid, CreateLoginLengthValidation(cfg))
	return LoginAndSpecification{valid}
}

func CreateLoginAcceptableValuesValidation(cfg config.Config) LoginAcceptableValuesValidation {
	return LoginAcceptableValuesValidation{cfg.Valid.LoginValidate.NonAcceptableValues}
}

func CreateLoginLengthValidation(cfg config.Config) LoginLengthValidation {
	return LoginLengthValidation{cfg.Valid.LoginValidate.MinLength,
		cfg.Valid.LoginValidate.MaxLength}
}

func CreatePhoneNumberSpecification(cfg config.Config) PhoneNumberSpecification {
	var valid []Specification[string]
	valid = append(valid, CreatePhoneNumberOrSpecification(cfg))
	return PhoneNumberSpecification{valid}
}

func CreatePhoneNumberOrSpecification(cfg config.Config) PhoneNumberAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreatePhoneNumberRussianRequiredValueValidation(cfg))
	return PhoneNumberAndSpecification{valid}
}

func CreatePhoneNumberRussianRequiredValueValidation(cfg config.Config) PhoneNumberRussianRequiredValueValidation {
	return PhoneNumberRussianRequiredValueValidation{}
}

func CreateFirstNameSpecification(cfg config.Config) FirstNameSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateFirstNameAndSpecification(cfg))
	return FirstNameSpecification{valid}
}

func CreateFirstNameAndSpecification(cfg config.Config) FirstNameAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateFirstNameLengthValidation(cfg))
	return FirstNameAndSpecification{valid}
}

func CreateFirstNameLengthValidation(cfg config.Config) FirstNameLengthValidation {
	return FirstNameLengthValidation{cfg.Valid.FirstNameValidate.MinLength, cfg.Valid.FirstNameValidate.MaxLength}
}

func CreateSecondNameSpecification(cfg config.Config) SecondNameSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateSecondNameAndSpecification(cfg))
	return SecondNameSpecification{valid}
}

func CreateSecondNameAndSpecification(cfg config.Config) SecondNameAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateSecondNameLengthValidation(cfg))
	return SecondNameAndSpecification{valid}
}

func CreateSecondNameLengthValidation(cfg config.Config) SecondNameLengthValidation {
	return SecondNameLengthValidation{cfg.Valid.SecondNameValidate.MinLength, cfg.Valid.SecondNameValidate.MaxLength}
}

func CreatePatronimicSpecification(cfg config.Config) PatronimicSpecification {
	var valid []Specification[string]
	valid = append(valid, CreatePatronimicAndSpecification(cfg))
	return PatronimicSpecification{valid}
}

func CreatePatronimicAndSpecification(cfg config.Config) PatronimicAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreatePatronimicLengthValidation(cfg))
	return PatronimicAndSpecification{valid}
}

func CreatePatronimicLengthValidation(cfg config.Config) PatronimicLengthValidation {
	return PatronimicLengthValidation{cfg.Valid.PatronimicValidate.MinLength, cfg.Valid.PatronimicValidate.MaxLength}
}

func CreateAvatarSpecification(cfg config.Config) AvatarSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateAvatarAndSpecification(cfg))
	return AvatarSpecification{valid}
}

func CreateAvatarAndSpecification(cfg config.Config) AvatarAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateAvatarLengthValidation(cfg))
	return AvatarAndSpecification{valid}
}

func CreateAvatarLengthValidation(cfg config.Config) AvatarWeightValidation {
	return AvatarWeightValidation{cfg.Valid.AvatarValidate.Weigth}
}
