# XDB

Database migration management using [Atlas](https://atlasgo.io/) with GORM schema loader.

## Prerequisites

- Go 1.21+
- Docker (for Atlas dev database)

## Install Atlas

```bash
# macOS
brew install ariga/tap/atlas

# Linux
curl -sSf https://atlasgo.sh | sh

# Or using Go
go install ariga.io/atlas/cmd/atlas@latest
```

Verify installation:

```bash
atlas version
```

## Create Migration File

Generate a new migration file based on your GORM models:

```bash
atlas migrate diff <migration_name> --env gorm
```

**Example:**

```bash
atlas migrate diff add_user_is_active_column --env gorm
```

This will:

1. Load schema from GORM models via `./loader/main.go`
2. Compare with existing migrations in `migrations/`
3. Generate new `.up.sql` and `.down.sql` files

## Run Migrations

Apply pending migrations to your database:

```bash
atlas migrate apply --env gorm --url "postgres://user:pass@localhost:5432/dbname?sslmode=disable"
```

**Using environment variable:**

```bash
export DATABASE_URL="postgres://user:pass@localhost:5432/dbname?sslmode=disable"
atlas migrate apply --env gorm --url "$DATABASE_URL"
```

## Other Useful Commands

```bash
# Check migration status
atlas migrate status --env gorm --url "$DATABASE_URL"

# Validate migrations
atlas migrate validate --env gorm

# View migration hash
atlas migrate hash --env gorm
```

## Project Structure

```
xdb/
├── atlas.hcl          # Atlas configuration
├── loader/main.go     # GORM schema loader
├── models/            # GORM model definitions
└── migrations/        # Generated SQL migration files
```
