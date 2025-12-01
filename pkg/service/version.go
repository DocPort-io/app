package service

import (
	"app/pkg/dto"
	"app/pkg/model"
	"context"

	"gorm.io/gorm"
)

type VersionService struct {
	db          *gorm.DB
	fileService *FileService
}

func NewVersionService(db *gorm.DB, fileService *FileService) *VersionService {
	return &VersionService{db: db, fileService: fileService}
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

func (s *VersionService) UploadFileToVersion(ctx context.Context, id string, uploadFileToVersionDto dto.UploadFileToVersionDto) (*model.File, error) {
	file, err := s.fileService.CreateFile(ctx, dto.CreateFileDto{
		Name: uploadFileToVersionDto.Name,
		Size: uploadFileToVersionDto.Size,
		Path: uploadFileToVersionDto.Path,
	})
	if err != nil {
		return nil, err
	}

	version, err := s.FindVersionById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.db.Model(&version).Association("Files").Append(file)
	if err != nil {
		return nil, err
	}

	return file, nil
}
