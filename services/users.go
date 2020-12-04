package services

import (
	"github.com/ruyjfs/example-golang/config"
	"github.com/ruyjfs/example-golang/graphql/model"
	"github.com/ruyjfs/example-golang/models"
)

type Users struct {
}

func (s *Users) GetAll(input *model.SearchUser) ([]*model.User, error) {
	// model := models.User
	db := config.Db()
	var users []*model.User
	db.Where(input).Find(&users)
	return users, nil
}

func (s *Users) Save(user *models.User) (*models.User, error) {
	config.Db().Create(&user)
	return user, nil
}
