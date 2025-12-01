package model

type Version struct {
	BaseModel
	Name        string
	Description string
	ProjectId   uint
	Files       []File `gorm:"many2many:version_files;"`
}
