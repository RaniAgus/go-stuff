package main

import (
	"context"
	"log"
	"net/http"

	"github.com/RaniAgus/go-starter/api"
	"github.com/RaniAgus/go-starter/data"
	"github.com/RaniAgus/go-starter/data/sqlc"
)

func main() {
	db := data.Connect()
	defer db.Close(context.Background())

	r := api.NewRouter(sqlc.New(db))

	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
