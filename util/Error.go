package util

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func NewErrorResponse(message string, err string) ErrorResponse {
	logrus.Error(message)

	return ErrorResponse{
		Message: message,
		Errors:  map[string]string{"error": err},
	}
}

func NewValidationErrorResponse(err error) ErrorResponse {
	res := ErrorResponse{
		Message: "Validation failed",
	}

	res.Errors = make(map[string]string)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			res.Errors[fieldError.Field()] = fieldError.Tag()
		}
	} else {
		res.Errors["error"] = err.Error()
	}

	return res
}
