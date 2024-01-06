package main

import (
	"context"
	"log"
	"reflect"

	db "github.com/RaniAgus/go-sqlc/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	url := "postgres://postgres:postgres@localhost:5432/authors"

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	query := db.New(conn)

	// list all authors
	authors, err := query.ListAuthors(ctx)
	log.Println(authors)

	// create an author
	inserted, err := query.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Agustin",
		Bio:  pgtype.Text{String: "un tipazo", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(inserted)

	// get the author we just inserted
	fetched, err := query.GetAuthor(ctx, inserted.ID)
	if err != nil {
		return err
	}
	log.Println(fetched)

	// prints true because both are equal
	log.Println(reflect.DeepEqual(inserted, fetched))

	return nil
}
