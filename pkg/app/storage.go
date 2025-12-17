package app

import (
	"app/pkg/storage"
	"fmt"
	"log"
)

func NewFileStorage(cfg *Config) storage.FileStorage {
	backend := storage.Type(cfg.Storage.Provider)
	if backend == storage.TypeFileSystem {
		fileStorage, err := storage.NewFilesystemStorage(cfg.Storage.Path)
		if err != nil {
			log.Fatalf("failed to initialize file storage backend %s: %s\n", backend, err)
		}

		return fileStorage
	}

	if backend == storage.TypeS3 {
		//TODO implement me
		panic("implement me")
	}

	log.Fatalf("failed to initialize file storage backend %s: %s\n", backend, fmt.Errorf("unsupported backend"))
	return nil
}
