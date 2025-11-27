package model

type Version struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   uint   `json:"projectId"`
	Files       []File `json:"files" gorm:"many2many:version_files;"`
}
