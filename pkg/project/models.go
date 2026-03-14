package project

import "time"

type Project struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Slug      string
	Name      string
}

type CreateProjectRequest struct {
	Slug string
	Name string
}

type UpdateProjectRequest struct {
	Slug string
	Name string
}
