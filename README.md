# go-starter

Golang starter project.

## Stack

- [task](https://taskfile.dev) for task management
- [goose](https://pressly.github.io/goose/) for database migrations
- [pgx](https://github.com/jackc/pgx) as a PostgreSQL driver
- [sqlc](https://sqlc.dev/) for generating code from SQL queries
- [chi](https://go-chi.io/) for routing
- [templ](https://templ.guide/) as a template engine
- [validator](https://pkg.go.dev/github.com/go-playground/validator/v10) for validating input data
- [uuid](https://pkg.go.dev/github.com/google/uuid) for generating UUIDs


## Install dev dependencies

```bash
chmod +x ./scripts/install-deps.sh
./scripts/install-deps.sh
```

## Create migration file

```bash
task migration name=<migration_name>
```

## Start server in watch mode

```bash
task
```
