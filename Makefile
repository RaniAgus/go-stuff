all: templ run

run:
	@go run cmd/go-templ/main.go

templ:
	@templ generate
