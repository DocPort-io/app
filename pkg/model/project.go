package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Slug     string    `json:"slug"`
	Name     string    `json:"name"`
	Versions []Version `json:"versions"`
}
