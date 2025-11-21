package postgresql

import (
	"fmt"
	"log"
	"time"

	"github.com/alirezazahiri/gotasks/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	DB *gorm.DB
}

// New creates a new database connection using the provided configuration
func New(cfg config.Config) (*Repository, error) {
	dsn := buildDSN(cfg.Repository.Postgres)

	gormConfig := &gorm.Config{
		Logger: getLogger(cfg.Env),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Repository{DB: db}, nil
}

// buildDSN constructs the PostgreSQL Data Source Name from config
func buildDSN(cfg config.PostgresConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
	)
}

// getLogger returns the appropriate GORM logger based on environment
func getLogger(env string) logger.Interface {
	if env == "production" {
		return logger.Default.LogMode(logger.Error)
	}
	return logger.Default.LogMode(logger.Info)
}

// Close closes the database connection
func (d *Repository) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	log.Println("Repository connection closed")
	return nil
}

// Ping checks if the database connection is alive
func (d *Repository) Ping() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection is alive")

	return nil
}
