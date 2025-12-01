package model

type Project struct {
	BaseModel
	Slug       string
	Name       string
	LocationId uint
	Location   Location
	Versions   []Version
}
