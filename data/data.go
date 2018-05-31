package data

import (
	"gopkg.in/go-playground/validator.v9"
)

type Message struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Body  string `json:"body" validate:"required"`
}

type MessageCreationResult struct {
	ID int `json:"id"`
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type ValidationResult struct {
	Valid   bool              `json:"valid"`
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors"`
}

func Validate(data interface{}) ValidationResult {
	raw := validator.New().Struct(data)

	if raw == nil {
		return ValidationResult{
			Valid: true,
		}
	}

	errors := raw.(validator.ValidationErrors)
	builtErrors := make([]ValidationError, len(errors))

	for i, e := range errors {
		builtErrors[i] = ValidationError{
			Field:  e.Field(),
			Reason: e.Tag(),
		}
	}

	return ValidationResult{
		Valid:  false,
		Errors: builtErrors,
	}
}
