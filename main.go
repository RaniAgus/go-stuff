package main

import (
	"context"
	"log"
	"net/http"

	"github.com/RaniAgus/go-templ-playground/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func withLayout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), templates.SidebarKey, templates.SidebarUser)
		ctx = context.WithValue(ctx, templates.HeaderKey, templates.HeaderAll)
		ctx = context.WithValue(ctx, templates.RouteKey, r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	cars := []templates.Car{
		{
			Make:     "Tesla",
			Model:    "Model Y",
			Price:    64950,
			Electric: true,
			Id:       1,
		},
		{
			Make:     "Ford",
			Model:    "F-Series",
			Price:    33850,
			Electric: false,
			Id:       2,
		},
		{
			Make:     "Toyota",
			Model:    "Corolla",
			Price:    29600,
			Electric: false,
			Id:       3,
		},
	}
	templates.Index(cars).Render(r.Context(), w)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.NoCache)
	r.Use(withLayout)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/cats", http.StatusMovedPermanently)
	})
	r.Get("/cats", handle)
	r.Get("/dogs", handle)
	r.Get("/dragons", handle)
	r.Get("/spiders", handle)
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	port := "3000"
	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
