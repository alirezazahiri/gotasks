package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alirezazahiri/gotasks/internal/config"
	"github.com/alirezazahiri/gotasks/internal/pkg/migrate"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	var (
		configPath     = flag.String("config", "config.yml", "path to config file")
		migrationsPath = flag.String("migrations", "migrations", "path to migrations directory")
		command        = flag.String("cmd", "up", "migration command: up, down, steps, version, force")
		steps          = flag.Int("steps", 1, "number of steps for steps command")
		version        = flag.Int("version", 0, "version number for force command")
	)
	flag.Parse()

	cfg := config.Load(*configPath)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Repository.Postgres.Host,
		cfg.Repository.Postgres.Username,
		cfg.Repository.Postgres.Password,
		cfg.Repository.Postgres.DBName,
		cfg.Repository.Postgres.Port,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	migrator, err := migrate.New(db, *migrationsPath)
	if err != nil {
		log.Fatalf("failed to create migrator: %v", err)
	}
	defer migrator.Close()

	switch *command {
	case "up":
		log.Println("Running migrations up...")
		if err := migrator.Up(); err != nil {
			log.Fatalf("failed to run migrations: %v", err)
		}
		log.Println("Migrations completed successfully")

	case "down":
		log.Println("Rolling back all migrations...")
		if err := migrator.Down(); err != nil {
			log.Fatalf("failed to rollback migrations: %v", err)
		}
		log.Println("Rollback completed successfully")

	case "steps":
		log.Printf("Running %d migration steps...\n", *steps)
		if err := migrator.Steps(*steps); err != nil {
			log.Fatalf("failed to run migration steps: %v", err)
		}
		log.Println("Migration steps completed successfully")

	case "version":
		v, dirty, err := migrator.Version()
		if err != nil {
			log.Fatalf("failed to get migration version: %v", err)
		}
		if dirty {
			log.Printf("Current version: %d (dirty)\n", v)
		} else {
			log.Printf("Current version: %d\n", v)
		}

	case "force":
		log.Printf("Forcing version to %d...\n", *version)
		if err := migrator.Force(*version); err != nil {
			log.Fatalf("failed to force version: %v", err)
		}
		log.Println("Version forced successfully")

	default:
		log.Fatalf("unknown command: %s", *command)
		os.Exit(1)
	}
}
