include .env

export POSTGRES_DB
export POSTGRES_USER
export POSTGRES_PASSWORD

export DBSTRING=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)

# Start local development server

all: goose sqlc templ tailwind run

goose: args ?= up
goose:
	@echo "=== Running migrations... ==="
	@goose -dir data/migrations postgres "$(DBSTRING)" $(args)

sqlc: args ?= generate
sqlc:
	@echo "=== Compiling queries... ==="
	@sqlc $(args)

templ: args ?= generate
templ:
	@echo "=== Compiling templates... ==="
	@templ $(args)

tailwind: args ?= --minify
tailwind:
	@echo "=== Minifying CSS... ==="
	@tailwindcss -i ./web/static/css/tailwind.css -o ./web/static/css/tailwind.min.css $(args)

run:
	@echo "=== Starting development server... ==="
	@go run cmd/server/main.go

# Common development tasks

seed:
	@echo "=== Seeding database... ==="
	@go run cmd/seeder/main.go

clean:
	@rm -rfv data/sqlc
	@find . -type f -name '*_templ.go' | xargs --no-run-if-empty rm -rfv
	@rm -rfv web/static/css/tailwind.min.css

watch-tailwind:
	@make --no-print-directory tailwind args=--watch

migration: name ?= $(error name is not set)
migration:
	@goose -dir data/migrations postgres "$(DBSTRING)" create $(name) sql

# Install dependencies

install:
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	sudo mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

.PHONY: all goose sqlc templ tailwind run clean watch migration install
