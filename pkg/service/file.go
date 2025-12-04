package service

import (
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/storage"
	"context"
	"os"
	"path"

	"github.com/google/uuid"
)

type FileService struct {
	queries     *database.Queries
	fileStorage storage.FileStorage
}

func NewFileService(queries *database.Queries, fileStorage storage.FileStorage) *FileService {
	return &FileService{queries: queries, fileStorage: fileStorage}
}

func (s *FileService) FindAllFiles(ctx context.Context, versionId *int64) ([]*database.File, error) {
	if versionId != nil {
		files, err := s.queries.ListFilesByVersionId(ctx, *versionId)
		if err != nil {
			return nil, err
		}
		return files, nil
	}

	files, err := s.queries.ListFiles(ctx)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (s *FileService) FindFileById(ctx context.Context, id *int64) (*database.File, error) {
	file, err := s.queries.GetFile(ctx, *id)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *FileService) CreateFile(ctx context.Context, createFileDto dto.CreateFileDto) (*database.File, error) {
	assetPath := buildFileAssetPath("")

	src, err := os.Open(createFileDto.Path)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	err = s.fileStorage.Save(ctx, assetPath, src)
	if err != nil {
		return nil, err
	}

	file, err := s.queries.CreateFile(ctx, &database.CreateFileParams{
		Name: createFileDto.Name,
		Size: createFileDto.Size,
		Path: assetPath,
	})
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *FileService) DeleteFile(ctx context.Context, id *int64) error {
	file, err := s.FindFileById(ctx, id)
	if err != nil {
		return err
	}

	err = s.queries.DeleteFile(ctx, file.ID)
	if err != nil {
		return err
	}

	err = s.fileStorage.Delete(ctx, file.Path)
	if err != nil {
		return err
	}

	return nil
}

func buildFileAssetPath(fileUuid string) string {
	if fileUuid == "" {
		fileUuid = uuid.NewString()
	}

	return path.Join("files", fileUuid)
}
