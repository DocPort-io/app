package model

type File struct {
	BaseModel
	Name string `json:"fileName"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}
