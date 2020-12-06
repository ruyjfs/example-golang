package core

import (
	"github.com/ruyjfs/example-golang/config"
	"gorm.io/gorm"
)

type Resolver struct {
}

func (r *Resolver) Db() *gorm.DB {
	return config.Db()
}
