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
	FindAllFiles(ctx context.Context, params *dto.FindAllFilesParams) (*dto.FindAllFilesResult, error)
	FindFileById(ctx context.Context, params *dto.FindFileByIdParams) (*dto.FindFileByIdResult, error)
	CreateFile(ctx context.Context, params *dto.CreateFileParams) (*dto.CreateFileResult, error)
	UploadFile(ctx context.Context, params *dto.UploadFileParams) (*dto.UploadFileResult, error)
	DownloadFile(ctx context.Context, params *dto.DownloadFileParams) (*dto.DownloadFileResult, error)
	DeleteFile(ctx context.Context, params *dto.DeleteFileParams) error
}

type fileServiceImpl struct {
	queries     *database.Queries
	fileStorage storage.FileStorage
}

func NewFileService(queries *database.Queries, fileStorage storage.FileStorage) FileService {
	return &fileServiceImpl{queries: queries, fileStorage: fileStorage}
}

func (s *fileServiceImpl) FindAllFiles(ctx context.Context, params *dto.FindAllFilesParams) (*dto.FindAllFilesResult, error) {
	files, err := s.queries.ListFilesByVersionId(ctx, params.VersionID)
	if err != nil {
		return nil, err
	}

	count, err := s.queries.CountFilesByVersionId(ctx, params.VersionID)
	if err != nil {
		return nil, err
	}

	return &dto.FindAllFilesResult{Files: files, Total: count}, nil
}

func (s *fileServiceImpl) FindFileById(ctx context.Context, params *dto.FindFileByIdParams) (*dto.FindFileByIdResult, error) {
	file, err := s.queries.GetFile(ctx, params.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &dto.FindFileByIdResult{File: file}, nil
}

func (s *fileServiceImpl) CreateFile(ctx context.Context, params *dto.CreateFileParams) (*dto.CreateFileResult, error) {
	file, err := s.queries.CreateFile(ctx, params.Name)
	if err != nil {
		return nil, err
	}

	return &dto.CreateFileResult{File: file}, nil
}

func (s *fileServiceImpl) UploadFile(ctx context.Context, params *dto.UploadFileParams) (*dto.UploadFileResult, error) {
	defer func(File multipart.File) {
		err := File.Close()
		if err != nil {
			log.Printf("error closing uploaded file: %v", err)
		}
	}(params.File)

	file, err := s.queries.GetFile(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if file.IsComplete {
		return nil, ErrFileAlreadyExists
	}

	assetPath := buildFileAssetPath("")

	mimeType, err := mimetype.DetectReader(params.File)
	if err != nil {
		return nil, err
	}

	_, err = params.File.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	err = s.fileStorage.Save(ctx, assetPath, params.File)
	if err != nil {
		return nil, err
	}

	mimeTypeString := mimeType.String()

	file, err = s.queries.UpdateFileWithUploadedFile(ctx, &database.UpdateFileWithUploadedFileParams{
		ID:       params.ID,
		Size:     &params.FileHeader.Size,
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

	return &dto.UploadFileResult{File: file}, nil
}

func (s *fileServiceImpl) DownloadFile(ctx context.Context, params *dto.DownloadFileParams) (*dto.DownloadFileResult, error) {
	findRes, err := s.FindFileById(ctx, &dto.FindFileByIdParams{ID: params.ID})
	if err != nil {
		return nil, err
	}
	file := findRes.File

	if !file.IsComplete {
		return nil, ErrIncompleteFile
	}

	reader, err := s.fileStorage.Retrieve(ctx, *file.Path)
	if err != nil {
		return nil, err
	}

	return &dto.DownloadFileResult{File: file, Reader: reader}, nil
}

func (s *fileServiceImpl) DeleteFile(ctx context.Context, params *dto.DeleteFileParams) error {
	findRes, err := s.FindFileById(ctx, &dto.FindFileByIdParams{ID: params.ID})
	if err != nil {
		return err
	}
	file := findRes.File

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
