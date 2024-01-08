package main

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/RaniAgus/go-starter/api"
	"github.com/RaniAgus/go-starter/data"
	"github.com/RaniAgus/go-starter/data/sqlc"
	"github.com/RaniAgus/go-starter/web"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := data.Connect()
	defer db.Close(context.Background())

	h := web.Handler{
		DB:       sqlc.New(db),
		Validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	h.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	r := api.NewRouter(h)

	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
