package models

import (
	"github.com/ruyjfs/example-golang/core"
)

type Perfil struct {
	Name        string `json:"name" gorm:"index;not null"`
	Description string `json:"description" gorm:"not null"`
	core.Model
}
