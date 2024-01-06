include .env

export POSTGRES_DB
export POSTGRES_USER
export POSTGRES_PASSWORD

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable

all: goose sqlc run

clean:
	@rm -rf db/sqlc

run:
	@go run cmd/go-sqlc/main.go

sqlc:
	@sqlc generate

migration: name ?= $(error name is not set)
migration:
	@goose -dir db/migrations create $(name) sql

goose:
	@goose -dir db/migrations up
