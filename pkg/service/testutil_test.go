package service

import (
	appRoot "app"
	"app/pkg/database"
	"database/sql"
	"log"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "modernc.org/sqlite"
)

// newTestDBAndQueries creates an in-memory sqlite database, runs embedded migrations,
// and returns the sql DB and sqlc Queries instance.
func newTestDBAndQueries(t *testing.T) (*sql.DB, *database.Queries) {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open sqlite in-memory db: %v", err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		t.Fatalf("failed to create sqlite driver: %v", err)
	}

	iofsDriver, err := iofs.New(appRoot.MigrationsFS, "migrations")
	if err != nil {
		t.Fatalf("failed to create migrations source: %v", err)
	}

	migrations, err := migrate.NewWithInstance("iofs", iofsDriver, "sqlite", driver)
	if err != nil {
		t.Fatalf("failed to init migrations: %v", err)
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		t.Fatalf("failed to run migrations: %v", err)
	}

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close test db: %v", err)
		}
	})

	return db, database.New(db)
}
