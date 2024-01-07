package web

import (
	"net/http"

	"github.com/RaniAgus/go-starter/data/sqlc"
	"github.com/RaniAgus/go-starter/web/templates"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	DB       sqlc.Querier
	Validate *validator.Validate
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h Handler) GetHome(w http.ResponseWriter, r *http.Request) error {
	return templates.ShowHome().Render(r.Context(), w)
}
