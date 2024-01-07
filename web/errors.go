package web

import "net/http"

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

func (h *Handler) HandlePageError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
