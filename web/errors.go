package web

import (
	"log"
	"net/http"

	"github.com/RaniAgus/go-starter/web/templates"
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
