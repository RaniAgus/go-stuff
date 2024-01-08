package web

import (
	"log"
	"net/http"
	"strings"

	"github.com/RaniAgus/go-starter/web/templates"
	"github.com/go-playground/validator/v10"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

func (err APIError) Error() string {
	return err.Message
}

func (h Handler) HandlePageError(w http.ResponseWriter, r *http.Request, err error) {
	msg := "Something went wrong."
	if apiError, ok := err.(APIError); ok {
		msg = apiError.Message
	} else {
		log.Println(err)
	}

	templates.ShowErrorPage(msg).Render(r.Context(), w)
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
