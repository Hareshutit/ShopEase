package domain

func CreateSpecificationManager() SpecificationManager {
	return SpecificationManager{
		Email:       CreateEmailSpecification(),
		Login:       CreateLoginSpecification(),
		PhoneNumber: CreatePhoneNumberSpecification(),
		Password:    CreatePassSpecification(),
		Name:        CreateNameSpecification(),
		Avatar:      CreateAvatarSpecification(),
	}
}

func CreatePassSpecification() PassSpecification {
	var valid []Specification[string]
	valid = append(valid, CreatePassAndSpecification())
	return PassSpecification{valid}
}

func CreatePassAndSpecification() PassAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreatePassLengthValidation(), CreatePassForRuneValidation())
	return PassAndSpecification{valid}
}

func CreatePassLengthValidation() PassLengthValidation {
	return PassLengthValidation{maxLength: 20, minLength: 8}
}

func CreatePassForRuneValidation() PassForRuneValidation {
	SpecialChar := make(map[rune]bool)
	SpecialChar['/'] = true
	SpecialChar['.'] = true
	SpecialChar[','] = true
	validspecialCharacters := SpecialChar
	validRune := make(map[string]Specification[rune])
	validRune["SC"] = passSpecialCharactersValidation{validspecialCharacters}
	validRune["UC"] = passUpperCaseValidation{}
	validRune["LC"] = passLowerCaseValidation{}
	return PassForRuneValidation{validRune}
}

func CreateEmailSpecification() PassSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateEmailAndSpecification())
	return PassSpecification{valid}
}

func CreateEmailAndSpecification() PassAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateEmailRequiredValueValidation())
	return PassAndSpecification{valid}
}

func CreateEmailRequiredValueValidation() EmailRequiredValueValidation {
	return EmailRequiredValueValidation{}
}

func CreateLoginSpecification() LoginSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateLoginAndSpecification())
	return LoginSpecification{valid}
}

func CreateLoginAndSpecification() LoginAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateLoginAcceptableValuesValidation())
	valid = append(valid, CreateLoginLengthValidation())
	return LoginAndSpecification{valid}
}

func CreateLoginAcceptableValuesValidation() LoginAcceptableValuesValidation {
	NonAcceptableValues := make(map[rune]bool)
	NonAcceptableValues['/'] = true
	NonAcceptableValues['.'] = true
	NonAcceptableValues[','] = true
	NonAcceptableValues['@'] = true
	return LoginAcceptableValuesValidation{NonAcceptableValues}
}

func CreateLoginLengthValidation() LoginLengthValidation {
	return LoginLengthValidation{minLength: 8, maxLength: 20}
}

func CreatePhoneNumberSpecification() PhoneNumberSpecification {
	var valid []Specification[string]
	valid = append(valid, CreatePhoneNumberOrSpecification())
	return PhoneNumberSpecification{valid}
}

func CreatePhoneNumberOrSpecification() PhoneNumberAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreatePhoneNumberRussianRequiredValueValidation())
	return PhoneNumberAndSpecification{valid}
}

func CreatePhoneNumberRussianRequiredValueValidation() PhoneNumberRussianRequiredValueValidation {
	return PhoneNumberRussianRequiredValueValidation{}
}

func CreateNameSpecification() NameSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateNameAndSpecification())
	return NameSpecification{valid}
}

func CreateNameAndSpecification() NameAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateNameLengthValidation())
	return NameAndSpecification{valid}
}

func CreateNameLengthValidation() NameLengthValidation {
	return NameLengthValidation{1, 15}
}

func CreateAvatarSpecification() AvatarSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateAvatarAndSpecification())
	return AvatarSpecification{valid}
}

func CreateAvatarAndSpecification() AvatarAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateAvatarLengthValidation())
	return AvatarAndSpecification{valid}
}

func CreateAvatarLengthValidation() AvatarWeightValidation {
	return AvatarWeightValidation{20971520}
}
