package domain

import (
	"encoding/base64"
	"regexp"
	"unicode"
)

type SpecificationManager struct {
	Email       Specification[string]
	Login       Specification[string]
	PhoneNumber Specification[string]
	Password    Specification[string]
	FirstName   Specification[string]
	SecondName  Specification[string]
	Patronimic  Specification[string]
	Avatar      Specification[string]
}

type typeSpecification interface {
	string | rune
}

type Specification[T typeSpecification] interface {
	IsValid(verifiable T) error
}

type PassSpecification struct {
	specifications []Specification[string]
}

func (e PassSpecification) IsValid(pass string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(pass); err != nil {
			return err
		}
	}
	return nil
}

type PassAndSpecification struct {
	specifications []Specification[string]
}

func (e PassAndSpecification) IsValid(pass string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(pass); err != nil {
			return err
		}
	}
	return nil
}

type PassLengthValidation struct {
	maxLength uint
	minLength uint
}

func (e PassLengthValidation) IsValid(pass string) error {
	if e.minLength < uint(len(pass)) && e.maxLength < uint(len(pass)) {
		return PassMaxLengthErr{e.maxLength}
	}
	if e.minLength > uint(len(pass)) && e.maxLength > uint(len(pass)) {
		return PassMinLengthErr{e.minLength}
	}
	return nil
}

type PassForRuneValidation struct {
	specifications map[string]Specification[rune]
}

func (e PassForRuneValidation) IsValid(pass string) error {
	errSC := e.specifications["SC"].IsValid(rune(pass[0]))
	errUC := e.specifications["UC"].IsValid(rune(pass[0]))
	errLC := e.specifications["LC"].IsValid(rune(pass[0]))
	for _, char := range pass[1:] {
		if errSC != nil {
			errSC = e.specifications["SC"].IsValid(char)
		}
		if errUC != nil {
			errUC = e.specifications["UC"].IsValid(char)
		}
		if errLC != nil {
			errLC = e.specifications["LC"].IsValid(char)
		}
	}
	if errSC != nil {
		return errSC
	}
	if errUC != nil {
		return errUC
	}
	if errLC != nil {
		return errLC
	}
	return nil
}

type passSpecialCharactersValidation struct {
	specialChar map[rune]bool
}

func (e passSpecialCharactersValidation) IsValid(char rune) error {
	if e.specialChar[char] {
		return nil
	}
	return PassSpecialCharactersErr{e.specialChar}
}

type passUpperCaseValidation struct{}

func (e passUpperCaseValidation) IsValid(char rune) error {
	if unicode.IsUpper(char) {
		return nil
	}
	return PassUpperCaseErr{}
}

type passLowerCaseValidation struct{}

func (e passLowerCaseValidation) IsValid(char rune) error {
	if unicode.IsLower(char) {
		return nil
	}
	return PassLowerCaseErr{}
}

type EmailSpecification struct {
	specifications []Specification[string]
}

func (e EmailSpecification) IsValid(email string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(email); err != nil {
			return err
		}
	}
	return nil
}

type EmailAndSpecification struct {
	specifications []Specification[string]
}

func (e EmailAndSpecification) IsValid(email string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(email); err != nil {
			return err
		}
	}
	return nil
}

type EmailRequiredValueValidation struct{}

func (e EmailRequiredValueValidation) IsValid(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return EmailRequiredValueErr{}
	} else {
		return nil
	}
}

type PhoneNumberSpecification struct {
	specifications []Specification[string]
}

func (e PhoneNumberSpecification) IsValid(number string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(number); err != nil {
			return err
		}
	}
	return nil
}

type PhoneNumberOrSpecification struct {
	specifications []Specification[string]
}

func (e PhoneNumberOrSpecification) IsValid(number string) error {
	var err error
	for _, specification := range e.specifications {
		if err = specification.IsValid(number); err == nil {
			return nil
		}
	}
	return err
}

type PhoneNumberAndSpecification struct {
	specifications []Specification[string]
}

func (e PhoneNumberAndSpecification) IsValid(number string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(number); err != nil {
			return err
		}
	}
	return nil
}

type PhoneNumberRussianRequiredValueValidation struct{}

func (e PhoneNumberRussianRequiredValueValidation) IsValid(phone string) error {
	emailRegex := regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`)
	if !emailRegex.MatchString(phone) {
		return EmailRequiredValueErr{}
	} else {
		return nil
	}
}

type LoginSpecification struct {
	specifications []Specification[string]
}

func (e LoginSpecification) IsValid(login string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(login); err != nil {
			return err
		}
	}
	return nil
}

type LoginAndSpecification struct {
	specifications []Specification[string]
}

func (e LoginAndSpecification) IsValid(login string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(login); err != nil {
			return err
		}
	}
	return nil
}

type LoginAcceptableValuesValidation struct {
	NonAcceptableValues map[rune]bool
}

func (e LoginAcceptableValuesValidation) IsValid(email string) error {
	for char := range e.NonAcceptableValues {
		if e.NonAcceptableValues[char] {
			return nil
		}
	}
	return LoginAcceptableValuesErr{e.NonAcceptableValues}
}

type LoginLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e LoginLengthValidation) IsValid(login string) error {
	if e.minLength < uint(len(login)) && e.maxLength < uint(len(login)) {
		return LoginMaxLengthErr{e.maxLength}
	}
	if e.minLength < uint(len(login)) && e.maxLength < uint(len(login)) {
		return LoginMinLengthErr{e.minLength}
	}
	return nil
}

type FirstNameSpecification struct {
	specifications []Specification[string]
}

func (e FirstNameSpecification) IsValid(name string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(name); err != nil {
			return err
		}
	}
	return nil
}

type FirstNameAndSpecification struct {
	specifications []Specification[string]
}

func (e FirstNameAndSpecification) IsValid(name string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(name); err != nil {
			return err
		}
	}
	return nil
}

type FirstNameLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e FirstNameLengthValidation) IsValid(name string) error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return FirstNameMaxLengthErr{e.maxLength}
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return FirstNameMinLengthErr{e.minLength}
	}
	return nil
}

type SecondNameSpecification struct {
	specifications []Specification[string]
}

func (e SecondNameSpecification) IsValid(name string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(name); err != nil {
			return err
		}
	}
	return nil
}

type SecondNameAndSpecification struct {
	specifications []Specification[string]
}

func (e SecondNameAndSpecification) IsValid(name string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(name); err != nil {
			return err
		}
	}
	return nil
}

type SecondNameLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e SecondNameLengthValidation) IsValid(name string) error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return SecondNameMaxLengthErr{e.maxLength}
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return SecondNameMinLengthErr{e.minLength}
	}
	return nil
}

type PatronimicSpecification struct {
	specifications []Specification[string]
}

func (e PatronimicSpecification) IsValid(name string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(name); err != nil {
			return err
		}
	}
	return nil
}

type PatronimicAndSpecification struct {
	specifications []Specification[string]
}

func (e PatronimicAndSpecification) IsValid(name string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(name); err != nil {
			return err
		}
	}
	return nil
}

type PatronimicLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e PatronimicLengthValidation) IsValid(name string) error {
	if e.minLength <= uint(len(name)) && e.maxLength <= uint(len(name)) {
		return PatronimicMaxLengthErr{e.maxLength}
	}
	if e.minLength <= uint(len(name)) && e.maxLength <= uint(len(name)) {
		return PatronimicMinLengthErr{e.minLength}
	}
	return nil
}

type AvatarSpecification struct {
	specifications []Specification[string]
}

func (e AvatarSpecification) IsValid(name string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(name); err != nil {
			return err
		}
	}
	return nil
}

type AvatarAndSpecification struct {
	specifications []Specification[string]
}

func (e AvatarAndSpecification) IsValid(name string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(name); err != nil {
			return err
		}
	}
	return nil
}

type AvatarWeightValidation struct {
	maxImageSize uint
}

func (e AvatarWeightValidation) IsValid(avatar string) error {
	if avatar == "" {
		return nil
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(avatar)))
	n, err := base64.StdEncoding.Decode(dst, []byte(avatar))
	if err != nil {
		return err
	}
	if n > int(e.maxImageSize) {
		return AvatarWeightErr{e.maxImageSize}
	}
	return nil
}
