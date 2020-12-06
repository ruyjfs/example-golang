package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ruyjfs/example-golang/graphql/generated"
	"github.com/ruyjfs/example-golang/graphql/model"
	"github.com/ruyjfs/example-golang/graphql/resolvers"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return new(resolvers.User).Create(input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	return new(resolvers.User).Update(input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input *model.DeleteInput) (*model.User, error) {
	return new(resolvers.User).Delete(input.ID)
}

func (r *queryResolver) Users(ctx context.Context, input *model.SearchUser) ([]*model.User, error) {
	return new(resolvers.User).All(input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
