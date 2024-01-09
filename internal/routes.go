package internal

import (
	"log"
	"net/http"

	"github.com/RaniAgus/go-starter/internal/handler"
	"github.com/RaniAgus/go-starter/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Serve(h handler.Handler) {
	r := chi.NewRouter()
	fs := http.FileServer(http.Dir("static"))

	r.Use(middleware.Logger)
	r.Get("/", route(h.GetHome, h.HandlePageError))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.NotFound(route(h.NotFound, h.HandlePageError))

	port := util.Getenv("PORT", "3000")

	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func route(handlerFn handler.HandlerFunc, errorFn handler.ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFn(w, r); err != nil {
			errorFn(w, r, err)
		}
	}
}
