//go:generate sqlc generate -f ../../sqlc.yaml
package app

import (
	"context"
	//"app/pkg/model"
	"database/sql"
	"log"

	//"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"

	"app/pkg/database"
	//"gorm.io/driver/sqlite"
	//"github.com/glebarez/sqlite"
	//"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	log.Fatalf("failed to connect database: %s\n", err)
	//}
	//
	//err = db.AutoMigrate(&model.User{}, &model.Project{}, &model.Version{}, &model.Location{}, &model.File{})
	//if err != nil {
	//	log.Fatalf("failed to migrate database: %s\n", err)
	//}

	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatalf("failed to open database: %s\n", err)
	}

	queries := database.New(db)

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("failed to begin transaction: %s\n", err)
	}
	defer tx.Rollback()

	queries.WithTx(tx)

	projects, err := queries.ListProjects(context.Background())
	if err != nil {
		log.Fatalf("failed to list projects: %s\n", err)
	}

	log.Printf("projects: %+v\n", projects)

	err = tx.Commit()
	if err != nil {
		log.Fatalf("failed to commit transaction: %s\n", err)
	}

	return nil
}
