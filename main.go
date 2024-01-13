package main

import (
	"log"
	"net/http"

	"github.com/RaniAgus/go-templ-playground/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.NoCache)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Index().Render(r.Context(), w)
	})
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	port := "3000"
	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
