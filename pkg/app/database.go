//go:generate sqlc generate -f ../../sqlc.yaml
package app

import (
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"

	"app/pkg/database"
)

func NewDatabase() (*sql.DB, *database.Queries) {
	db, err := sql.Open("sqlite", "file:test.db?_pragma=foreign_keys(ON)")
	if err != nil {
		log.Fatalf("failed to open database: %s\n", err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatalf("failed to create database driver: %s\n", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance("file://./migrations", "sqlite", driver)
	if err != nil {
		log.Fatalf("failed to create migration source: %s\n", err)
	}

	err = migrations.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Println("no migrations applied, database schema is up to date")
	} else if err != nil {
		log.Fatalf("failed to run migrations: %s\n", err)
	}

	queries := database.New(db)

	return db, queries
}
