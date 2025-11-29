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

func (s *FileService) FindAllFiles(ctx context.Context, versionId string) ([]model.File, error) {
	var files []model.File
	var err error

	if versionId != "" {
		version, err := gorm.G[model.Version](s.db).Preload("Files", nil).Where("id = ?", versionId).First(ctx)
		if err != nil {
			return nil, err
		}

		files = version.Files

		return files, err
	}

	files, err = gorm.G[model.File](s.db).Find(ctx)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (s *FileService) FindFileById(ctx context.Context, id string) (*model.File, error) {
	file, err := gorm.G[model.File](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return &file, nil
}
