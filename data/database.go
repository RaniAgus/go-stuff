package data

import (
	"context"
	"fmt"
	"log"

	"github.com/RaniAgus/go-starter/util"
	"github.com/jackc/pgx/v5"
)

func Connect() *pgx.Conn {
	ctx := context.Background()
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		util.Getenv("POSTGRES_USER"),
		util.Getenv("POSTGRES_PASSWORD"),
		util.Getenv("POSTGRES_HOST", "localhost"),
		util.Getenv("POSTGRES_PORT", "5432"),
		util.Getenv("POSTGRES_DB"),
	)

	db, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
