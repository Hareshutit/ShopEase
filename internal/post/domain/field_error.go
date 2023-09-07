package domain

import (
	"fmt"
)

type DescriptionMinLengthErr struct {
	length uint
}

func (e DescriptionMinLengthErr) Error() string {
	return fmt.Sprintf("Min length of Description: %d\n", e.length)
}

type DescriptionMaxLengthErr struct {
	length uint
}

func (e DescriptionMaxLengthErr) Error() string {
	return fmt.Sprintf("Max length of Desctription: %d\n", e.length)
}

type TitleMinLengthErr struct {
	length uint
}

func (e TitleMinLengthErr) Error() string {
	return fmt.Sprintf("Max length of Title: %d\n", e.length)
}

type TitleMaxLengthErr struct {
	length uint
}

func (e TitleMaxLengthErr) Error() string {
	return fmt.Sprintf("Max length of Title: %d\n", e.length)
}

type ImagesMaxLength struct {
	length uint
}

func (e ImagesMaxLength) Error() string {
	return fmt.Sprintf("Max number of Images: %d\n", e.length)
}

type ImagesMinLength struct {
	length uint
}

func (e ImagesMinLength) Error() string {
	return fmt.Sprintf("Min number of Images: %d\n", e.length)
}

type ImageWeightErr struct {
	weight uint
}

func (e ImageWeightErr) Error() string {
	return fmt.Sprintf("Max weight of Image: %d\n", e.weight)
}

type TagsMaxLength struct {
	length uint
}

func (e TagsMaxLength) Error() string {
	return fmt.Sprintf("Max length of Tags: %d\n", e.length)
}

type TagsMinLength struct {
	length uint
}

func (e TagsMinLength) Error() string {
	return fmt.Sprintf("Min length of Tags: %d\n", e.length)
}
