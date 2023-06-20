# poster
## Requirements (Must be installed)
- Go v1.20.4
- Go Migrate CLI (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- Postgres

## How to setup DB:
1. Install Posgres
2. Create a new `poster` database in Posgres
```CREATE DATABASE poster;```

## How to run DB migrations
Requires a `.env` file in the same project directory with the following properties:
- `DB_HOST`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `DB_PORT`
- `DB_SSLMODE`
- `DB_TIMEZONE`

### Windows
1. `cd` into the project directory
2. Run `init.cmd` for every new terminal

**Create Migration Files**

To create migration files with `MIGRATION_NAME`:

```$ .\cmig.cmd MIGRATION_NAME```

**Migrate Up**

- To migrate up to the latest version:
```$ .\mig.cmd up```

- To migrate up by `X` steps from the current version:
```$ .\mig.cmd up X```

**Migrate Down**

- To migrate down to undo all migrations:
```$ .\mig.cmd down```

- To migrate down by `X` steps from the current version:
```$ .\mig.cmd down X```

**Force Migrate Step**

To force a migration to step `X`:
```$ .\mig.cmd force X```

### mac/Linux
1. `cd` into the project directory

**Create Migration Files**

To create migration files with `MIGRATION_NAME`:

```$ ./cmig.sh MIGRATION_NAME```

**Migrate Up**

- To migrate up to the latest version:
```$ ./mig.sh up```

- To migrate up by `X` steps from the current version:
```$ ./mig.sh up X```

**Migrate Down**

- To migrate down to undo all migrations:
```$ ./mig.sh down```

- To migrate down by `X` steps from the current version:
```$ ./mig.sh down X```

**Force Migrate Step**

To force a migration to step `X`:
```$ ./mig.sh force X```

For more info on migration, visit the [golang-migrate](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md) page

## Note on DB migration files

1. Add the PSQL statements in the new `.up` files in `database/migration`
2. Add the inverse of the PSQL statements in the `.down` files in `database/migration`

## How to run:
```go run .```
