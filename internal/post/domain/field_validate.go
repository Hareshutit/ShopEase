package domain

import (
	"encoding/base64"
)

type SpecificationManager struct {
	Title      Specification[string]
	Desciption Specification[string]
	Images     Specification[string]
	Tags       Specification[string]
}

type typeSpecification interface {
	string | rune
}

type Specification[T typeSpecification] interface {
	IsValid(verifiable T) error
}

type TitleSpecification struct {
	specifications []Specification[string]
}

func (e TitleSpecification) IsValid(title string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(title); err != nil {
			return err
		}
	}
	return nil
}

type TitleAndSpecification struct {
	specifications []Specification[string]
}

func (e TitleAndSpecification) IsValid(title string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(title); err != nil {
			return err
		}
	}
	return nil
}

type TitleLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e TitleLengthValidation) IsValid(title string) error {
	if e.minLength < uint(len(title)) && e.maxLength < uint(len(title)) {
		return TitleMaxLengthErr{e.maxLength}
	}
	if e.minLength > uint(len(title)) && e.maxLength > uint(len(title)) {
		return TitleMinLengthErr{e.minLength}
	}
	return nil
}

type DescriptionSpecification struct {
	specifications []Specification[string]
}

func (e DescriptionSpecification) IsValid(desc string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(desc); err != nil {
			return err
		}
	}
	return nil
}

type DescriptionAndSpecification struct {
	specifications []Specification[string]
}

func (e DescriptionAndSpecification) IsValid(desc string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(desc); err != nil {
			return err
		}
	}
	return nil
}

type DescriptionLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e DescriptionLengthValidation) IsValid(desc string) error {
	if e.minLength <= uint(len(desc)) && e.maxLength <= uint(len(desc)) {
		return DescriptionMaxLengthErr{e.maxLength}
	}
	if e.minLength >= uint(len(desc)) && e.maxLength >= uint(len(desc)) {
		return DescriptionMinLengthErr{e.minLength}
	}
	return nil
}

type ImagesSpecification struct {
	specifications []Specification[string]
}

func (e ImagesSpecification) IsValid(image string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(image); err != nil {
			return err
		}
	}
	return nil
}

type ImagesAndSpecification struct {
	specifications []Specification[string]
}

func (e ImagesAndSpecification) IsValid(image string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(image); err != nil {
			return err
		}
	}
	return nil
}

type ImagesWeightValidation struct {
	minImages    uint
	maxImages    uint
	maxImageSize uint
}

func (e ImagesWeightValidation) IsValid(image string) error {
	if image == "" {
		return nil
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(image)))
	n, err := base64.StdEncoding.Decode(dst, []byte(image))
	if err != nil {
		return err
	}
	if n > int(e.maxImageSize) {
		return ImageWeightErr{e.maxImageSize}
	}
	return nil
}

type TagsSpecification struct {
	specifications []Specification[string]
}

func (e TagsSpecification) IsValid(tags string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(tags); err != nil {
			return err
		}
	}
	return nil
}

type TagsAndSpecification struct {
	specifications []Specification[string]
}

func (e TagsAndSpecification) IsValid(tags string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(tags); err != nil {
			return err
		}
	}
	return nil
}

type TagsLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e TagsLengthValidation) IsValid(tags string) error {
	if e.minLength <= uint(len(tags)) && e.maxLength <= uint(len(tags)) {
		return DescriptionMaxLengthErr{e.maxLength}
	}
	if e.minLength >= uint(len(tags)) && e.maxLength >= uint(len(tags)) {
		return DescriptionMinLengthErr{e.minLength}
	}
	return nil
}
