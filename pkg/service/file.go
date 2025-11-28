package service

import (
	"app/pkg/dto"
	"app/pkg/model"
	"app/pkg/storage"
	"context"
	"os"
	"path"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileService struct {
	db          *gorm.DB
	fileStorage storage.FileStorage
}

func NewFileService(db *gorm.DB, fileStorage storage.FileStorage) *FileService {
	return &FileService{db: db, fileStorage: fileStorage}
}

func (s *FileService) CreateFile(ctx context.Context, createFileDto dto.CreateFileDto) (*model.File, error) {
	assetPath := path.Join("files", uuid.NewString())

	src, err := os.Open(createFileDto.Path)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	err = s.fileStorage.Save(ctx, assetPath, src)
	if err != nil {
		return nil, err
	}

	var file = &model.File{
		Name: createFileDto.Name,
		Size: createFileDto.Size,
		Path: assetPath,
	}

	err = gorm.G[model.File](s.db).Create(ctx, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}
