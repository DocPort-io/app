package app

import (
	"app/pkg/storage"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func NewFileStorage(backend storage.Type) storage.FileStorage {
	if backend == storage.TypeFileSystem {
		fileStorage, err := storage.NewFilesystemStorage(viper.GetString("storage.path"))
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
