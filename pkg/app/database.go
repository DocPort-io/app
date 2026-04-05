package app

import (
	"app/pkg/platform/config"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	migratePgx "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	// Import for migration source "file" registration
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"app/pkg/database"
)

func NewDatabase(config config.DatabaseConfig) *database.Queries {
	ctx := context.Background()

	sqlDB, err := sql.Open("pgx", config.DSN)
	if err != nil {
		log.Fatalf("failed to open database: %s\n", err)
	}

	driver, err := migratePgx.WithInstance(sqlDB, &migratePgx.Config{})
	if err != nil {
		log.Fatalf("failed to create database driver: %s\n", err)
	}

	iofsDriver, err := iofs.New(database.Migrations, "migrations")
	if err != nil {
		log.Fatalf("failed to create migration source: %s\n", err)
	}

	migrations, err := migrate.NewWithInstance("iofs", iofsDriver, "pgx/v5", driver)
	if err != nil {
		log.Fatalf("failed to create migration source: %s\n", err)
	}

	err = migrations.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations applied, database schema is up to date")
		} else {
			log.Fatalf("failed to run migrations: %s\n", err)
		}
	} else {
		log.Println("migrations applied, database schema has been updated")
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("failed to close database connection: %s\n", err)
	}

	pool, err := pgxpool.New(ctx, config.DSN)
	if err != nil {
		log.Fatalf("failed to create database connection pool: %s\n", err)
	}

	queries := database.New(pool)

	return queries
}
