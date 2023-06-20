# poster
## Requirements (Must be installed)
- Go v1.20.4
- Go Migrate CLI (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- Postgres

## How to setup DB:
1. Install Posgres
2. Create a new `poster` database in Posgres
>`CREATE DATABASE poster;`

## How to run DB migrations
Requires a `.env` file in the same `/poster` directory with the following properties:
- `DB_HOST`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `DB_PORT`
- `DB_SSLMODE`
- `DB_TIMEZONE`

### Windows
1. Open the terminal (cmd)
2. Run `init.cmd` for every new terminal

**Create Migration Files**

To create migration files with `MIGRATION_NAME`, run:
>`.\createmig.cmd MIGRATION_NAME`

**Migrate Up**

- To migrate up to the latest version, run:
>`.\mig.cmd up`

- To migrate up by `X` steps from the current version, run:
>`.\mig.cmd up X`

**Migrate Down**

- To migrate down to undo all migrations, run:
>`.\mig.cmd down`

- To migrate down by `X` steps from the current version, run:
>`.\mig.cmd down X`

**Force Migrate Step**

To force a migration to step `X`, run:
>`.\mig.cmd force X`

### mac/Linux
1. Create an export using environment values (run once for new terminal)

`export POSTER_PSQL_URL='postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}'`

2. Run migration (choose either up or down)

`migrate -database ${POSTER_PSQL_URL} -path database/migration {up|down}`

More info: https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

## How to create DB migration files
1. Create the up and down migration files

`migrate create -ext sql -dir database/migration -seq MIGRATION_NAME_HERE`

2. Add the PSQL statements in the new `.up` files in `database/migration`
3. Add the inverse of the PSQL statements in the `.down` files in `database/migration`

## How to run:
`go run .`
