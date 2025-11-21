# Database Migrations

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for database schema management.

## Quick Start

### Windows

```bash
# Run all pending migrations
migrate.bat up

# Rollback all migrations
migrate.bat down

# Check current migration version
migrate.bat version

# Run specific number of migration steps
migrate.bat steps 1      # forward 1 step
migrate.bat steps -1     # backward 1 step

# Force database to specific version (use with caution)
migrate.bat force 1

# Create new migration files
migrate.bat create add_users_table
```

### Linux/macOS

```bash
# Run all pending migrations
make migrate-up

# Rollback all migrations
make migrate-down

# Check current migration version
make migrate-version

# Run specific number of migration steps
make migrate-steps N=1   # forward 1 step
make migrate-steps N=-1  # backward 1 step

# Force database to specific version (use with caution)
make migrate-force V=1

# Create new migration files
make migrate-create NAME=add_users_table
```

## Migration File Structure

Each migration consists of two files:
- `{version}_{name}.up.sql` - Applied when migrating up
- `{version}_{name}.down.sql` - Applied when migrating down

Example:
```
migrations/
  000001_init.up.sql
  000001_init.down.sql
  000002_add_users.up.sql
  000002_add_users.down.sql
```

## Best Practices

1. **Always create both up and down migrations** - Ensure rollback is possible
2. **Test migrations locally** before applying to production
3. **Keep migrations small** - One logical change per migration
4. **Never modify applied migrations** - Create new ones instead
5. **Use transactions when possible** - Most DDL statements in PostgreSQL are transactional
6. **Handle failures gracefully** - If migration fails, database is marked as "dirty"

## Handling Dirty State

If a migration fails midway, the database enters a "dirty" state. To recover:

```bash
# Check current version and dirty state
migrate.bat version  # or make migrate-version

# Fix the issue, then force the version
migrate.bat force 1  # or make migrate-force V=1

# Try running migrations again
migrate.bat up  # or make migrate-up
```

## Configuration

Migrations read database configuration from `config.yml`:

```yaml
repository: 
  postgres:
    username: postgres
    password: postgres
    host: localhost
    port: 5432
    dbname: gotasks
```

## Programmatic Usage

You can also use the migration package in your Go code:

```go
import (
    "github.com/alirezazahiri/gotasks/internal/pkg/migrate"
)

migrator, err := migrate.New(db, "migrations")
if err != nil {
    log.Fatal(err)
}
defer migrator.Close()

// Run all pending migrations
if err := migrator.Up(); err != nil {
    log.Fatal(err)
}
```

