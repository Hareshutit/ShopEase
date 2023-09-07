package domain

import (
	"fmt"
	"strconv"
)

type PassMinLengthErr struct {
	length uint
}

func (e PassMinLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина пароля: %d\n", e.length)
}

type PassMaxLengthErr struct {
	length uint
}

func (e PassMaxLengthErr) Error() string {
	return fmt.Sprintf("Максимальная длина пароля: %d\n", e.length)
}

type PassSpecialCharactersErr struct {
	specialChar map[rune]bool
}

func (e PassSpecialCharactersErr) Error() string {
	var str string
	for char := range e.specialChar {
		str = fmt.Sprintf(str, strconv.QuoteRune(char), ", ")
	}
	str = str[0 : len(str)-3]
	return fmt.Sprintf("Пароль не содержит специальных символов из данного списка: %s\n", str)
}

type PassUpperCaseErr struct{}

func (e PassUpperCaseErr) Error() string {
	return "Пароль не содержит заглавные буквы"
}

type PassLowerCaseErr struct{}

func (e PassLowerCaseErr) Error() string {
	return "Пароль не содержит строчные буквы"
}

type EmailRequiredValueErr struct{}

func (e EmailRequiredValueErr) Error() string {
	return "Введите валидный email"
}

type EmailMaxLengthErr struct {
	length uint
}

func (e EmailMaxLengthErr) Error() string {
	return fmt.Sprintf("Максимальная длина email: %d\n", e.length)
}

type EmailMinLengthErr struct {
	length uint
}

func (e EmailMinLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина email: %d\n", e.length)
}

type PhoneNumberRequiredValueErr struct{}

func (e PhoneNumberRequiredValueErr) Error() string {
	return "Введите валидный телефонный номер"
}

type PhoneNumberMaxLengthErr struct {
	length uint
}

func (e PhoneNumberMaxLengthErr) Error() string {
	return fmt.Sprintf("Максимальная длина телефонного номера: %d\n", e.length)
}

type PhoneNumberMinLengthErr struct {
	length uint
}

func (e PhoneNumberMinLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина телефонного номера: %d\n", e.length)
}

type LoginAcceptableValuesErr struct {
	nonAcceptableValues map[rune]bool
}

func (e LoginAcceptableValuesErr) Error() string {
	var str string
	for char := range e.nonAcceptableValues {
		str = fmt.Sprintf(str, strconv.QuoteRune(char), ", ")
	}
	str = str[0 : len(str)-3]
	return fmt.Sprintf("Логин содержит недопустимые символы: %s\n", str)
}

type LoginMaxLengthErr struct {
	length uint
}

func (e LoginMaxLengthErr) Error() string {
	return fmt.Sprintf("Максимальная длина логина: %d\n", e.length)
}

type LoginMinLengthErr struct {
	length uint
}

func (e LoginMinLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина логина: %d\n", e.length)
}

type FirstNameMinLengthErr struct {
	length uint
}

func (e FirstNameMinLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина имени: %d\n", e.length)
}

type FirstNameMaxLengthErr struct {
	length uint
}

func (e FirstNameMaxLengthErr) Error() string {
	return fmt.Sprintf("Максимальная длина имени: %d\n", e.length)
}

type SecondNameMinLengthErr struct {
	length uint
}

func (e SecondNameMinLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина фамилии: %d\n", e.length)
}

type SecondNameMaxLengthErr struct {
	length uint
}

func (e SecondNameMaxLengthErr) Error() string {
	return fmt.Sprintf("Максимальная длина фамилии: %d\n", e.length)
}

type PatronimicMinLengthErr struct {
	length uint
}

func (e PatronimicMinLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина отчества: %d\n", e.length)
}

type PatronimicMaxLengthErr struct {
	length uint
}

func (e PatronimicMaxLengthErr) Error() string {
	return fmt.Sprintf("Максимальная длина отчества: %d\n", e.length)
}

type CompanyTypeErr struct {
	typeComp map[string]bool
}

func (e CompanyTypeErr) Error() string {
	var str string
	for typeComp := range e.typeComp {
		str = fmt.Sprintf(str, typeComp, ", ")
	}
	str = str[0 : len(str)-3]
	return fmt.Sprintf("Неправильные тип компании. Введите допустимый из данного списка: %s\n", str)
}

type CompanyNameMinLengthErr struct {
	length uint
}

func (e CompanyNameMinLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина названия компании: %d\n", e.length)
}

type CompanyNameMaxLengthErr struct {
	length uint
}

func (e CompanyNameMaxLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина названия компании: %d\n", e.length)
}

type AdressMinLengthErr struct {
	length uint
}

func (e AdressMinLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина адреса: %d\n", e.length)
}

type AdressMaxLengthErr struct {
	length uint
}

func (e AdressMaxLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина адреса: %d\n", e.length)
}

type AvatarWeightErr struct {
	weight uint
}

func (e AvatarWeightErr) Error() string {
	return fmt.Sprintf("Вес изображения выше допустимого: %d\n", e.weight)
}
