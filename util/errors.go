package util

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err APIError) Error() string {
	return err.Message
}

func GetValidationErrorFields(err error) []string {
	errors := []string{}
	for _, err := range err.(validator.ValidationErrors) {
		// If the error is a slice or map, we remove the key from the field name
		arrIndex := strings.Index(err.Field(), "[")
		if arrIndex > -1 {
			errors = append(errors, err.Field()[:arrIndex])
		} else {
			errors = append(errors, err.Field())
		}
	}
	return errors
}
