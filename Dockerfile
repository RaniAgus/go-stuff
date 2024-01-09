FROM cgr.dev/chainguard/go AS builder

ENV GOBIN=/usr/local/bin

RUN go install github.com/go-task/task/v3/cmd/task@latest \
 && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest \
 && go install github.com/pressly/goose/v3/cmd/goose@latest \
 && go install github.com/a-h/templ/cmd/templ@latest

ADD https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 /usr/local/bin/tailwindcss

RUN chmod +x /usr/local/bin/tailwindcss

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN task build

FROM cgr.dev/chainguard/glibc-dynamic

WORKDIR /app

COPY --from=builder /app/bin/go-starter ./
COPY --from=builder /app/static ./static

CMD ["./go-starter"]
