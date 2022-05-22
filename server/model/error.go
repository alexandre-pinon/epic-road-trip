package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetValErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "e164":
		return "Invalid phone format"
	case "min":
		return fmt.Sprintf("Should be at least %s characters long", fe.Param())
	case "max":
		return fmt.Sprintf("Should be less than %s characters", fe.Param())
	case "len":
		return fmt.Sprintf("Should be exactly %s characters long", fe.Param())
	}
	return "Unknown error"
}

type AppError struct {
	StatusCode int
	Err        error
}

func (appError *AppError) Error() string {
	return appError.Err.Error()
}
