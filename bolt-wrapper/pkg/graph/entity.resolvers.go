package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bolt-wrapper/pkg/graph/generated"
	"bolt-wrapper/pkg/graph/model"
	"context"
)

func (r *entityResolver) FindUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return &model.User{
		Username: username,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
