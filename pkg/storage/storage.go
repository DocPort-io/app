package storage

import (
	"context"
	"io"
)

var (
	TypeFileSystem = "fs"
	TypeS3         = "s3"
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
