package main

import (
	"context"

	"github.com/RaniAgus/go-starter/data"
	"github.com/RaniAgus/go-starter/data/sqlc"
)

func main() {
	db := data.Connect()
	defer db.Close(context.Background())

	_ = sqlc.New(db)

	// Insert your seed data here
}
