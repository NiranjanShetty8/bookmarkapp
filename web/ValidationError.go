package web

import "fmt"

type ValidationError struct {
	ErrorKey string            `json:"errorKey"`
	Errors   map[string]string `json:"errors"`
}

func (err ValidationError) Error() string {
	return fmt.Sprintf("%s", err.Errors)
}

func NewValidationError(err string, failedValidation map[string]string) *ValidationError {
	return &ValidationError{
		ErrorKey: err,
		Errors:   failedValidation,
	}
}
