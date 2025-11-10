package model

import (
	"app/pkg/model/primitives"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Nickname string              `json:"nickname"`
	Address  string              `json:"address"`
	Location primitives.Location `json:"location"`
}
