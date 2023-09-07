package domain

import config "github.com/Hareshutit/ShopEase/config/post"

func CreateSpecificationManager(cfg config.Config) SpecificationManager {
	return SpecificationManager{
		Title:      CreateTitleSpecification(cfg),
		Desciption: CreateDescriptionSpecification(cfg),
		Images:     CreateImagesSpecification(cfg),
		Tags:       CreateTagsSpecification(cfg),
	}
}

func CreateTitleSpecification(cfg config.Config) TitleSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateTitleAndSpecification(cfg))
	return TitleSpecification{valid}
}

func CreateTitleAndSpecification(cfg config.Config) TitleAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateTitleLengthValidation(cfg))
	return TitleAndSpecification{valid}
}

func CreateTitleLengthValidation(cfg config.Config) TitleLengthValidation {
	return TitleLengthValidation{cfg.Valid.TitleValidate.MinLength, cfg.Valid.TitleValidate.MaxLength}
}

func CreateDescriptionSpecification(cfg config.Config) DescriptionSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateDescriptionAndSpecification(cfg))
	return DescriptionSpecification{valid}
}

func CreateDescriptionAndSpecification(cfg config.Config) DescriptionAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateDescriptionLengthValidation(cfg))
	return DescriptionAndSpecification{valid}
}

func CreateDescriptionLengthValidation(cfg config.Config) DescriptionLengthValidation {
	return DescriptionLengthValidation{cfg.Valid.DescriptionValidate.MinLength, cfg.Valid.DescriptionValidate.MaxLength}
}

func CreateImagesSpecification(cfg config.Config) ImagesSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateImagesAndSpecification(cfg))
	return ImagesSpecification{valid}
}

func CreateImagesAndSpecification(cfg config.Config) ImagesAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateImagesCountValidation(cfg))
	return ImagesAndSpecification{valid}
}

func CreateImagesCountValidation(cfg config.Config) ImagesWeightValidation {
	return ImagesWeightValidation{cfg.Valid.ImagesValidate.MinCount, cfg.Valid.ImagesValidate.MaxCount, cfg.Valid.ImagesValidate.Weigth}
}

func CreateTagsSpecification(cfg config.Config) TagsSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateTagsAndSpecidication(cfg))
	return TagsSpecification{valid}
}

func CreateTagsAndSpecidication(cfg config.Config) TagsAndSpecification {
	var valid []Specification[string]
	valid = append(valid, CreateTagsLengthValidation(cfg))
	return TagsAndSpecification{valid}
}

func CreateTagsLengthValidation(cfg config.Config) TagsLengthValidation {
	return TagsLengthValidation{cfg.Valid.TagsValidate.MinLength, cfg.Valid.TagsValidate.MaxLength}
}
