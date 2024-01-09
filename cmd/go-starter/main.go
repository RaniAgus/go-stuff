package main

import (
	"context"

	"github.com/RaniAgus/go-starter/internal"
	"github.com/RaniAgus/go-starter/internal/handler"
	"github.com/RaniAgus/go-starter/internal/sql"
)

func main() {
	db := internal.NewDatabase()
	defer db.Close(context.Background())

	h := handler.Handler{
		Queries:  sql.New(db),
		Validate: internal.NewValidator(),
	}

	internal.Serve(h)
}
