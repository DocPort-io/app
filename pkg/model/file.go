package model

type File struct {
	BaseModel
	Name     string
	Size     int64
	Path     string
	Versions []Version `gorm:"many2many:version_files;"`
}
