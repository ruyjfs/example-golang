package services

import (
	"gorm.io/gorm"

	"github.com/ruyjfs/example-golang/config"
	"github.com/ruyjfs/example-golang/core"
	"github.com/ruyjfs/example-golang/graphql/model"

	// "github.com/ruyjfs/example-golang/graphql/model"
	"github.com/ruyjfs/example-golang/models"
)

type Users struct {
	core.Service
}

func (s *Users) ById(id int) (*gorm.DB, model.User, error) {
	var model model.User
	result := s.Db().First(&model, id)
	return result, model, nil
}

func GetTest(search interface{}, model ...interface{}) *gorm.DB {
	db := config.Db()
	// var users []*models.User

	// var newdata []*models.User
	// var newdata2 []interface{}
	// db.Where(search).Find(&newdata)

	// var test *[]interface{}
	// title := "User"
	// var test2 = models["User"]{Name:}

	result := db.Where(search).Find(&model)
	// log.Println(&result)
	// model = model
	// db.Where(search).Find(&test2)

	// model <- &test
	// return models
	// log.Println(newdata[0], "<->", newdata2, "---", &model)
	// aa := result
	// log.Println("---", test2, title)
	// var test2 map[string]interface{}

	// test2 = test

	return result
}

func test1(data string) {
	data = "AHHHHHH"
}

func (s *Users) All(user interface{}) (*gorm.DB, []*models.User) {
	var list []*models.User
	result := s.Db().Where(user).Find(&list)
	return result, list
}

func (s *Users) Create(user *models.User) (*gorm.DB, *models.User) {
	result := s.Db().Create(&user)
	return result, user
}

func (s *Users) Update(user *models.User) (*gorm.DB, *models.User) {
	result := s.Db().Save(&user)
	return result, user
}

func (s *Users) Delete(id int) (*gorm.DB, models.User) {
	var model models.User
	s.Db().First(&model, id)
	result := s.Db().Delete(&model)
	return result, model
}
