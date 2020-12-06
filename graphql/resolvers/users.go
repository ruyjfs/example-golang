package resolvers

import (
	"github.com/ruyjfs/example-golang/core"
	"github.com/ruyjfs/example-golang/graphql/model"
	"github.com/ruyjfs/example-golang/models"
)

type User struct {
	core.Service
}

func (m User) All(input *model.SearchUser) ([]*model.User, error) {
	var list []*model.User
	result := m.Db().Where(input).Find(&list)
	return list, result.Error
}

func (m User) Create(input model.NewUser) (*model.User, error) {
	data := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	result := m.Db().Create(&data)
	dataReturn := &model.User{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
	return dataReturn, result.Error
}

func (m User) Update(input model.UpdateUser) (*model.User, error) {
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	result := m.Db().Save(&user)
	var userGraphQL = &model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return userGraphQL, result.Error

}
func (m User) Delete(id int) (*model.User, error) {
	var data models.User
	m.Db().First(&data, id)
	result := m.Db().Delete(&data)
	var userGraphQL = &model.User{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
	return userGraphQL, result.Error

}
