package web

import "fmt"

//Implements Error Interface
type ValidationError struct {
	ErrorKey string            `json:"errorKey"`
	Errors   map[string]string `json:"errors"`
}

//Returns Error String
func (err ValidationError) Error() string {
	return fmt.Sprintf("%s", err.Errors)
}

func NewValidationError(err string, failedValidation map[string]string) *ValidationError {
	return &ValidationError{
		ErrorKey: err,
		Errors:   failedValidation,
	}
}
