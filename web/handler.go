package web

import (
	"net/http"

	"github.com/RaniAgus/go-starter/data/sqlc"
	"github.com/RaniAgus/go-starter/web/templates"
)

type Handler struct {
	DB sqlc.Querier
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h *Handler) GetHome(w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, "/films", http.StatusSeeOther)
	return nil
}

func (h *Handler) GetFilms(w http.ResponseWriter, r *http.Request) error {
	films, err := h.DB.ListFilms(r.Context())
	if err != nil {
		return err
	}

	return templates.ShowFilms(films).Render(r.Context(), w)
}

func (h *Handler) PostFilm(w http.ResponseWriter, r *http.Request) error {
	film, err := h.DB.CreateFilm(r.Context(), sqlc.CreateFilmParams{
		Title:    r.PostFormValue("title"),
		Director: r.PostFormValue("director"),
	})
	if err != nil {
		return err
	}

	return templates.SwapNewFilm(film).Render(r.Context(), w)
}
