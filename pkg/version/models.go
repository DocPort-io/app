package version

import "time"

type Version struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description *string
	ProjectID   int64
}

type CreateVersionRequest struct {
	Name        string
	Description *string
	ProjectId   int64
}

type UpdateVersionRequest struct {
	Name        string
	Description *string
}

type AttachFileRequest struct {
	FileID int64
}

type DetachFileRequest struct {
	FileID int64
}
