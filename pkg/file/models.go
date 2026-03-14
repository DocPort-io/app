package file

import (
	"mime/multipart"
	"time"
)

type File struct {
	ID         int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Size       *int64
	Path       *string
	MimeType   *string
	IsComplete bool
}

type CreateFileRequest struct {
	Name string `json:"name" validate:"required" example:"example.pdf"`
}

type UploadFileRequest struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}
