package main

import (
	"log"

	"github.com/RaniAgus/go-templ/internal/routes"
)

func main() {
	app := routes.NewRouter()

	log.Fatal(app.Start(":8080"))
}
