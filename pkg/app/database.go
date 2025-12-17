package app

import (
	appRoot "app"
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"

	// Import for migration source "file" registration
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	// Import for database driver "sqlite" registration
	_ "modernc.org/sqlite"

	"app/pkg/database"
)

func NewDatabase(cfg *Config) *database.Queries {
	databaseDriver := cfg.Database.Driver
	databaseUrl := cfg.Database.URL

	db, err := sql.Open(databaseDriver, databaseUrl)
	if err != nil {
		log.Fatalf("failed to open database: %s\n", err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatalf("failed to create database driver: %s\n", err)
	}

	iofsDriver, err := iofs.New(appRoot.MigrationsFS, "migrations")
	if err != nil {
		log.Fatalf("failed to create migration source: %s\n", err)
	}

	migrations, err := migrate.NewWithInstance("iofs", iofsDriver, databaseDriver, driver)
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

	queries := database.New(db)

	return queries
}
