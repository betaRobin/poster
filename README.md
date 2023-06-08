# poster
## Requirements (Must be installed)
- Go v1.20.4
- Go Migrate CLI (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- Postgres

## How to setup DB:
1. Install Posgres
2. Create a new `poster` database in Posgres

`CREATE DATABASE poster;`

## How to run DB migrations
1. Create an export using environment values (run once for new terminal)

```export POSTER_PSQL_URL='postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}'```

2. Run migration (choose either up or down)

`migrate -database ${POSTER_PSQL_URL} -path database/migration {up|down}`

More info: https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

## How to create DB migration files
1. Create the up and down migration files

`migrate create -ext sql -dir database/migration -seq MIGRATION_NAME_HERE`

2. Add the PSQL statements in the new `.up` files in `database/migration`
3. Add the inverse of the PSQL statements in the `.down` files in `database/migration`

## How to run:
```go run .```
