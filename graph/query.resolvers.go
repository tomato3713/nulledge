package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/tomato3713/nulledge/graph/model"
)

// Page is the resolver for the page field.
func (r *queryResolver) Page(ctx context.Context, id string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented: Page - page"))
}

// Pages is the resolver for the pages field.
func (r *queryResolver) Pages(ctx context.Context) ([]*model.Page, error) {
	panic(fmt.Errorf("not implemented: Pages - pages"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }