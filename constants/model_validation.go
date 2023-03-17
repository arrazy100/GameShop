package constants

import "github.com/go-playground/validator/v10"

type ErrorStructValidation struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct[T any](s T) []*ErrorStructValidation {
	var errors []*ErrorStructValidation

	v := validator.New()
	err := v.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorStructValidation

			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, &element)
		}
	}

	return errors
}
