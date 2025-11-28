package service

import (
	"app/pkg/dto"
	"app/pkg/model"
	"app/pkg/storage"
	"context"
	"fmt"
	"io"
	"path"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VersionService struct {
	db          *gorm.DB
	fileStorage storage.FileStorage
}

func NewVersionService(db *gorm.DB, fileStorage storage.FileStorage) *VersionService {
	return &VersionService{db: db, fileStorage: fileStorage}
}

func (s *VersionService) FindAllVersions(ctx context.Context, projectId string) ([]model.Version, error) {
	var versions []model.Version
	var err error

	if projectId != "" {
		versions, err = gorm.G[model.Version](s.db).Where("project_id = ?", projectId).Find(ctx)
		if err != nil {
			return nil, err
		}
		return versions, nil
	}

	versions, err = gorm.G[model.Version](s.db).Find(ctx)
	if err != nil {
		return nil, err
	}
	return versions, nil
}

func (s *VersionService) FindVersionById(ctx context.Context, id string) (*model.Version, error) {
	version, err := gorm.G[model.Version](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return &version, nil
}

func (s *VersionService) CreateVersion(ctx context.Context, dto dto.CreateVersionDto) (*model.Version, error) {
	version := dto.ToModel()

	err := gorm.G[model.Version](s.db).Create(ctx, version)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (s *VersionService) UpdateVersion(ctx context.Context, id string, dto dto.UpdateVersionDto) (*model.Version, error) {
	version := dto.ToModel()

	rowsAffected, err := gorm.G[model.Version](s.db).Where("id = ?", id).Updates(ctx, *version)
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	updatedVersion, err := gorm.G[model.Version](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return &updatedVersion, nil
}

func (s *VersionService) DeleteVersion(ctx context.Context, id string) error {
	rowsAffected, err := gorm.G[model.Version](s.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (s *VersionService) UploadFileToVersion(ctx context.Context, id string, fileStream io.ReadSeeker, fileName string, size int64) (*model.File, error) {
	fileUuid := uuid.NewString()
	assetPath := path.Join("files", fmt.Sprintf("%s", fileUuid))

	err := s.fileStorage.Save(ctx, assetPath, fileStream)
	if err != nil {
		return nil, err
	}

	var file = &model.File{
		Name: fileName,
		Size: size,
		Path: assetPath,
	}

	err = gorm.G[model.File](s.db).Create(ctx, file)
	if err != nil {
		return nil, err
	}

	version, err := gorm.G[model.Version](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	err = s.db.Model(&version).Association("Files").Append(file)
	if err != nil {
		return nil, err
	}

	return file, nil
}
