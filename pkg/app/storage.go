package app

import (
	"app/pkg/storage"
	"fmt"
	"log"
)

func NewFileStorage(backend string) storage.FileStorage {
	if backend == storage.TypeFileSystem {
		fileStorage, err := storage.NewFilesystemStorage("./files")
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
