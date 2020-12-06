package core

import (
	"github.com/ruyjfs/example-golang/config"
	"gorm.io/gorm"
)

type Service struct {
}

func (m *Service) Db() *gorm.DB {
	return config.Db()
}

func (s *Service) Find(m interface{}) {

}

// func Min (type T Ordered) (a, b T) T {
// 	if a < b {
// 			return a
// 	}
// 	return b
// }

// type List[type T] struct {
// 	elems []T
// }

// contract Sequence(T) {
// 	T string, []byte
// }
