//go:generate sqlc generate -f ../../sqlc.yaml
package app

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	"app/pkg/database"
)

func NewDatabase() (*sql.DB, *database.Queries) {
	db, err := sql.Open("sqlite", "file:test.db?_pragma=foreign_keys(ON)")
	if err != nil {
		log.Fatalf("failed to open database: %s\n", err)
	}

	queries := database.New(db)

	return db, queries
}
