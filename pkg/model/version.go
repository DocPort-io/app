package model

type Version struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectID   uint   `json:"projectId"`
}
