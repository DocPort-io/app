package model

type Project struct {
	BaseModel
	Slug     string    `json:"slug"`
	Name     string    `json:"name"`
	Versions []Version `json:"versions"`
}
