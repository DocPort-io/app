package file

import (
	"app/pkg/storage"
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"path"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

var (
	ErrFileNotComplete     = errors.New("file not complete")
	ErrFileAlreadyComplete = errors.New("file already complete")
)

type Service interface {
	GetById(ctx context.Context, id int64) (File, error)
	List(ctx context.Context, versionId *int64, limit, offset int64) ([]File, error)
	Create(ctx context.Context, req CreateFileRequest) (File, error)
	UploadFile(ctx context.Context, id int64, req UploadFileRequest) (File, error)
	Download(ctx context.Context, id int64) (File, io.ReadSeekCloser, error)
	Delete(ctx context.Context, id int64) error
}

type service struct {
	repository  Repository
	fileStorage storage.FileStorage
}

func NewFileService(repository Repository, fileStorage storage.FileStorage) Service {
	return &service{repository: repository, fileStorage: fileStorage}
}

func (s *service) GetById(ctx context.Context, id int64) (File, error) {
	return s.repository.GetById(ctx, id)
}

func (s *service) List(ctx context.Context, versionId *int64, limit, offset int64) ([]File, error) {
	return s.repository.List(ctx, versionId, limit, offset)
}

func (s *service) Create(ctx context.Context, req CreateFileRequest) (File, error) {
	file := File{
		Name: req.Name,
	}
	return s.repository.Create(ctx, file)
}

func (s *service) UploadFile(ctx context.Context, id int64, req UploadFileRequest) (File, error) {
	defer func(File multipart.File) {
		err := File.Close()
		if err != nil {
			log.Printf("error closing uploaded file: %v", err)
		}
	}(req.File)

	file, err := s.repository.GetById(ctx, id)
	if err != nil {
		return File{}, err
	}

	if file.IsComplete {
		return File{}, ErrFileAlreadyComplete
	}

	assetPath := buildFileAssetPath(uuid.NewString())

	mimeType, err := mimetype.DetectReader(req.File)
	if err != nil {
		return File{}, err
	}

	_, err = req.File.Seek(0, io.SeekStart)
	if err != nil {
		return File{}, err
	}

	err = s.fileStorage.Save(ctx, assetPath, req.File)
	if err != nil {
		return File{}, err
	}

	mimeTypeString := mimeType.String()

	file.Size = &req.FileHeader.Size
	file.Path = &assetPath
	file.MimeType = &mimeTypeString
	file.IsComplete = true

	file, err = s.repository.Update(ctx, file)
	if err != nil {
		fileDeleteErr := s.fileStorage.Delete(ctx, assetPath)
		if fileDeleteErr != nil {
			log.Printf("error deleting file during upload: %v", fileDeleteErr)
		}
		return File{}, err
	}

	return file, nil
}

func (s *service) Download(ctx context.Context, id int64) (File, io.ReadSeekCloser, error) {
	file, err := s.repository.GetById(ctx, id)
	if err != nil {
		return File{}, nil, err
	}

	if !file.IsComplete {
		return File{}, nil, ErrFileNotComplete
	}

	reader, err := s.fileStorage.Retrieve(ctx, *file.Path)
	if err != nil {
		return File{}, nil, err
	}

	return file, reader, nil
}

func (s *service) Delete(ctx context.Context, id int64) error {
	file, err := s.repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	err = s.repository.Delete(ctx, id)
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
	return path.Join("files", fileUuid)
}
