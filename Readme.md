# ASSIGNMENT 2

## Development Environtment

- Go fiber
- Postgresql
- Go fiber session

## Features

- Sign Up
- Send email activation link
- Activate account
- Sign In
- Sign Out



## Run Locally

Clone the project

```bash
  git clone https://github.com/fernandojec/assignment-2
```

Migrate Database

```sql
    execute sql at directory docs/dbMigration/db_Migration.sql
```

Update setting in .env file


Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod tidy
```

Go to the cmd/web directory

```bash
  cd cmd/web/
```

Start the server

```bash
  go run main.go
```

