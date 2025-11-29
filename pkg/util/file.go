package util

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MultipartFile struct {
	Name string
	Size int64
	Path string
}

func SaveMultipartFileToTemp(ctx *gin.Context, name string) (multipartFile *MultipartFile, err error) {
	fileHeader, err := ctx.FormFile(name)
	if err != nil {
		return nil, err
	}

	tempDir, err := os.MkdirTemp("", "uploads")
	if err != nil {
		return nil, err
	}

	var tempFilePath = path.Join(tempDir, uuid.NewString())

	err = ctx.SaveUploadedFile(fileHeader, tempFilePath, 0700)
	if err != nil {
		return nil, err
	}

	return &MultipartFile{
		Name: fileHeader.Filename,
		Size: fileHeader.Size,
		Path: tempFilePath,
	}, nil
}
