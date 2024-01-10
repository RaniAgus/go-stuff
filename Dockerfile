FROM golang:1-alpine AS go-builder

ENV GOBIN=/usr/local/bin

RUN go install github.com/go-task/task/v3/cmd/task@latest \
 && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest \
 && go install github.com/pressly/goose/v3/cmd/goose@latest \
 && go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN task build


FROM node:20-alpine AS node-builder

WORKDIR /app/templates

COPY templates ./

WORKDIR /app/styles

COPY styles ./

RUN npm install

RUN npm run build


FROM cgr.dev/chainguard/glibc-dynamic

WORKDIR /app

COPY --from=go-builder /app/bin/go-starter ./
COPY --from=go-builder /app/public/ ./public
COPY --from=node-builder /app/public/ ./public

ENTRYPOINT ["./go-starter"]
