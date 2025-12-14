package storage

import (
	"context"
	"io"
)

type Type string

const (
	TypeFileSystem Type = "filesystem"
	TypeS3         Type = "s3"
)

type ObjectInfo struct {
	Path string
	Size int64
}

type WalkFunc func(info ObjectInfo) error

type FileStorage interface {
	Save(ctx context.Context, relativePath string, data io.Reader) error
	Retrieve(ctx context.Context, relativePath string) (io.ReadCloser, error)
	Delete(ctx context.Context, relativePath string) error
	List(ctx context.Context, root string) ([]ObjectInfo, error)
	Walk(ctx context.Context, root string, walkFunc WalkFunc) error
}
