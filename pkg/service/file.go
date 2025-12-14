package service

import (
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/storage"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
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

func (s *FileService) FindAllFiles(ctx context.Context, versionId int64) ([]*database.File, int64, error) {
	files, err := s.queries.ListFilesByVersionId(ctx, versionId)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.queries.CountFilesByVersionId(ctx, versionId)
	if err != nil {
		return nil, 0, err
	}

	return files, count, nil
}

func (s *FileService) FindFileById(ctx context.Context, id int64) (*database.File, error) {
	file, err := s.queries.GetFile(ctx, id)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *FileService) CreateFile(ctx context.Context, createFileDto *dto.CreateFileDto) (*database.File, error) {
	file, err := s.queries.CreateFile(ctx, createFileDto.Name)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *FileService) UploadFile(ctx context.Context, id int64, uploadFileDto *dto.UploadFileDto) (*database.File, error) {
	fileStream, err := uploadFileDto.FileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer func(fileStream multipart.File) {
		err := fileStream.Close()
		if err != nil {
			log.Printf("error closing file stream: %v", err)
		}
	}(fileStream)

	assetPath := buildFileAssetPath("")

	err = s.fileStorage.Save(ctx, assetPath, fileStream)
	if err != nil {
		return nil, err
	}

	file, err := s.queries.UpdateFileWithUploadedFile(ctx, &database.UpdateFileWithUploadedFileParams{
		ID:   id,
		Size: &uploadFileDto.FileHeader.Size,
		Path: &assetPath,
	})
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *FileService) DownloadFile(ctx context.Context, id int64) (*database.File, io.ReadCloser, error) {
	file, err := s.FindFileById(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	if file.Path == nil {
		return nil, nil, fmt.Errorf("file has no path")
	}

	reader, err := s.fileStorage.Retrieve(ctx, *file.Path)
	if err != nil {
		return nil, nil, err
	}

	return file, reader, nil
}

func (s *FileService) DeleteFile(ctx context.Context, id int64) error {
	file, err := s.FindFileById(ctx, id)
	if err != nil {
		return err
	}

	err = s.queries.DeleteFile(ctx, file.ID)
	if err != nil {
		return err
	}

	if file.Path != nil {
		err = s.fileStorage.Delete(ctx, *file.Path)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildFileAssetPath(fileUuid string) string {
	if fileUuid == "" {
		fileUuid = uuid.NewString()
	}

	return path.Join("files", fileUuid)
}
