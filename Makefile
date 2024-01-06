run:
	@go run main.go

tailwind:
	@tailwindcss -i ./assets/css/tailwind.css -o ./assets/css/tailwind.min.css $(TAILWIND_FLAGS)

tailwind-watch: TAILWIND_FLAGS = --watch
tailwind-watch: tailwind

tailwind-dev: TAILWIND_FLAGS = --watch
tailwind-dev: tailwind
