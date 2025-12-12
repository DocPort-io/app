package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type filesystemStorage struct {
	root             *os.Root
	absoluteRootPath string
}

func NewFilesystemStorage(rootPath string) (FileStorage, error) {
	if err := os.MkdirAll(rootPath, 0700); err != nil {
		return nil, fmt.Errorf("failed to create root directory '%s': %w", rootPath, err)
	}

	root, err := os.OpenRoot(rootPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open root directory '%s': %w", rootPath, err)
	}

	absoluteRootPath, err := filepath.Abs(rootPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path of root directory '%s': %w", rootPath, err)
	}

	return &filesystemStorage{
		root:             root,
		absoluteRootPath: absoluteRootPath,
	}, nil
}

func (s *filesystemStorage) Save(_ context.Context, relativePath string, data io.Reader) error {
	relativePath = filepath.FromSlash(relativePath)

	if err := s.root.MkdirAll(filepath.Dir(relativePath), 0o700); err != nil {
		return fmt.Errorf("failed to create directories for relativePath '%s': %w", relativePath, err)
	}

	tmpName := relativePath + "." + uuid.NewString() + ".tmp"

	tmpFile, err := s.root.Create(tmpName)
	if err != nil {
		return fmt.Errorf("failed to create temporary file '%s': %w", tmpName, err)
	}

	_, err = io.Copy(tmpFile, data)
	if err != nil {
		err := tmpFile.Close()
		if err != nil {
			log.Printf("failed to close file stream: %v", err)
		}
		_ = s.root.Remove(tmpName)
		return fmt.Errorf("failed to write data to temporary file '%s': %w", tmpName, err)
	}

	if err = tmpFile.Close(); err != nil {
		_ = s.root.Remove(tmpName)
		return fmt.Errorf("failed to close temporary file '%s': %w", tmpName, err)
	}

	if err = s.root.Rename(tmpName, relativePath); err != nil {
		_ = s.root.Remove(tmpName)
		return fmt.Errorf("failed to move temporary file '%s' to '%s': %w", tmpName, relativePath, err)
	}

	return nil
}

func (s *filesystemStorage) Retrieve(ctx context.Context, relativePath string) (io.ReadCloser, error) {
	relativePath = filepath.FromSlash(relativePath)

	file, err := s.root.Open(relativePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s': %w", relativePath, err)
	}

	return file, nil
}

func (s *filesystemStorage) Delete(ctx context.Context, relativePath string) error {
	//TODO implement me
	panic("implement me")
}

func (s *filesystemStorage) List(ctx context.Context, root string) ([]ObjectInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *filesystemStorage) Walk(ctx context.Context, root string, walkFunc WalkFunc) error {
	//TODO implement me
	panic("implement me")
}
