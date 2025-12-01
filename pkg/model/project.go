package model

type Project struct {
	BaseModel
	Slug     string
	Name     string
	Location Location
	Versions []Version
}
