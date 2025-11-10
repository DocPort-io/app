package model

import "gorm.io/gorm"

type Version struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectID   uint   `json:"projectId"`
}
