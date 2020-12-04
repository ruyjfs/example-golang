package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mholt/binding"
	"github.com/ruyjfs/example-golang/graphql/generated"
	"github.com/ruyjfs/example-golang/graphql/model"
	"github.com/ruyjfs/example-golang/models"
	"github.com/ruyjfs/example-golang/services"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user *models.User

	// user := models.User{
	// 	Name:     input.Name,
	// 	Email:    input.Email,
	// 	Password: input.Password,
	// }

	// new(services.Users).Save(&user)

	var userGraphQL = &model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	errs := binding.Bind(input, user)

	userGraphQL := model.User(input)

	return userGraphQL, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, input *model.SearchUser) ([]*model.User, error) {
	return new(services.Users).GetAll(input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
