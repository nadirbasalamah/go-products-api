package model

import "github.com/go-playground/validator/v10"

type UserInput struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func (userInput UserInput) ValidateStruct() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(userInput)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
	}

	return errors
}
