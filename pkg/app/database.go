package app

import (
	"app/pkg/model"
	"log"

	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s\n", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Project{}, &model.Version{}, &model.Location{}, &model.File{})
	if err != nil {
		log.Fatalf("failed to migrate database: %s\n", err)
	}

	return db
}
