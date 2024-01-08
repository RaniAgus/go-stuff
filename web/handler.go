package web

import (
	"net/http"
	"strings"

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
	http.Redirect(w, r, "/films", http.StatusSeeOther)
	return nil
}

func (h Handler) GetFilms(w http.ResponseWriter, r *http.Request) error {
	films, err := h.DB.ListFilms(r.Context())
	if err != nil {
		return err
	}

	return templates.ShowFilms(films).Render(r.Context(), w)
}

func (h Handler) PostFilm(w http.ResponseWriter, r *http.Request) error {
	form := templates.NewFilmForm{
		Title:    strings.Trim(r.PostFormValue("title"), " "),
		Director: strings.Fields(r.PostFormValue("director")), // debería usar varios inputs en vez de esto
	}

	err := h.Validate.Struct(form)
	if err != nil {
		return templates.ShowNewFilmForm(form, GetValidationErrorFields(err)).Render(r.Context(), w)
	}

	film, err := h.DB.CreateFilm(r.Context(), sqlc.CreateFilmParams{
		Title:    form.Title,
		Director: strings.Join(form.Director, " "),
	})
	if err != nil {
		return err
	}

	return templates.SwapNewFilm(film).Render(r.Context(), w)
}
