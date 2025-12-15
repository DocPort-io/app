package service

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/storage"
	"context"
	"database/sql"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"path"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

var (
	ErrFileAlreadyExists = errors.New("file already exists")
	ErrIncompleteFile    = errors.New("file not complete")
)

type FileService interface {
	FindAllFiles(ctx context.Context, versionId int64) ([]*database.File, int64, error)
	FindFileById(ctx context.Context, id int64) (*database.File, error)
	CreateFile(ctx context.Context, createFileDto *dto.CreateFileDto) (*database.File, error)
	UploadFile(ctx context.Context, id int64, uploadFileDto *dto.UploadFileDto) (*database.File, error)
	DownloadFile(ctx context.Context, id int64) (*database.File, io.ReadSeekCloser, error)
	DeleteFile(ctx context.Context, id int64) error
}

type fileServiceImpl struct {
	queries     *database.Queries
	fileStorage storage.FileStorage
}

func NewFileService(queries *database.Queries, fileStorage storage.FileStorage) FileService {
	return &fileServiceImpl{queries: queries, fileStorage: fileStorage}
}

func (s *fileServiceImpl) FindAllFiles(ctx context.Context, versionId int64) ([]*database.File, int64, error) {
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

func (s *fileServiceImpl) FindFileById(ctx context.Context, id int64) (*database.File, error) {
	file, err := s.queries.GetFile(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return file, nil
}

func (s *fileServiceImpl) CreateFile(ctx context.Context, createFileDto *dto.CreateFileDto) (*database.File, error) {
	file, err := s.queries.CreateFile(ctx, createFileDto.Name)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *fileServiceImpl) UploadFile(ctx context.Context, id int64, uploadFileDto *dto.UploadFileDto) (*database.File, error) {
	defer func(File multipart.File) {
		err := File.Close()
		if err != nil {
			log.Printf("error closing uploaded file: %v", err)
		}
	}(uploadFileDto.File)

	file, err := s.queries.GetFile(ctx, id)
	if err != nil {
		return nil, err
	}

	if file.IsComplete {
		return nil, ErrFileAlreadyExists
	}

	assetPath := buildFileAssetPath("")

	mimeType, err := mimetype.DetectReader(uploadFileDto.File)
	if err != nil {
		return nil, err
	}

	_, err = uploadFileDto.File.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	err = s.fileStorage.Save(ctx, assetPath, uploadFileDto.File)
	if err != nil {
		return nil, err
	}

	mimeTypeString := mimeType.String()

	file, err = s.queries.UpdateFileWithUploadedFile(ctx, &database.UpdateFileWithUploadedFileParams{
		ID:       id,
		Size:     &uploadFileDto.FileHeader.Size,
		Path:     &assetPath,
		MimeType: &mimeTypeString,
	})
	if err != nil {
		fileDeleteErr := s.fileStorage.Delete(ctx, assetPath)
		if fileDeleteErr != nil {
			log.Printf("error deleting file during upload: %v", fileDeleteErr)
		}
		return nil, err
	}

	return file, nil
}

func (s *fileServiceImpl) DownloadFile(ctx context.Context, id int64) (*database.File, io.ReadSeekCloser, error) {
	file, err := s.FindFileById(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	if !file.IsComplete {
		return nil, nil, ErrIncompleteFile
	}

	reader, err := s.fileStorage.Retrieve(ctx, *file.Path)
	if err != nil {
		return nil, nil, err
	}

	return file, reader, nil
}

func (s *fileServiceImpl) DeleteFile(ctx context.Context, id int64) error {
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
