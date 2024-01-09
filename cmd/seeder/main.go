package main

import (
	"context"

	"github.com/RaniAgus/go-starter/internal"
	"github.com/RaniAgus/go-starter/internal/sql"
)

func main() {
	db := internal.NewDatabase()
	defer db.Close(context.Background())

	_ = sql.New(db)

	// Insert your seed data here
}
