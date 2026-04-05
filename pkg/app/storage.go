package app

import (
	"app/pkg/platform/config"
	"app/pkg/storage"
	"fmt"
	"log"
)

func NewFileStorage(cfg config.StorageConfig) storage.FileStorage {
	backend := storage.Type(cfg.Provider)
	if backend == storage.TypeFileSystem {
		fileStorage, err := storage.NewFilesystemStorage(cfg.Path)
		if err != nil {
			log.Fatalf("failed to initialize file storage backend %s: %s\n", backend, err)
		}

		return fileStorage
	}

	log.Fatalf("failed to initialize file storage backend %s: %s\n", backend, fmt.Errorf("unsupported backend"))
	return nil
}
