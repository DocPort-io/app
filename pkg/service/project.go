package service

import (
	"app/pkg/dto"
	"app/pkg/model"
	"context"

	"gorm.io/gorm"
)

type ProjectService struct {
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{db: db}
}

func (s *ProjectService) FindAllProjects(ctx context.Context) ([]model.Project, error) {
	projects, err := gorm.G[model.Project](s.db).Find(ctx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *ProjectService) FindProjectById(ctx context.Context, id string) (*model.Project, error) {
	project, err := gorm.G[model.Project](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *ProjectService) CreateProject(ctx context.Context, dto dto.CreateProjectDto) (*model.Project, error) {
	project := dto.ToModel()

	err := gorm.G[model.Project](s.db).Create(ctx, project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) UpdateProject(ctx context.Context, id string, dto dto.UpdateProjectDto) (*model.Project, error) {
	project := dto.ToModel()

	rowsAffected, err := gorm.G[model.Project](s.db).Where("id = ?", id).Updates(ctx, *project)
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	updatedProject, err := gorm.G[model.Project](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return &updatedProject, nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, id string) error {
	rowsAffected, err := gorm.G[model.Project](s.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
