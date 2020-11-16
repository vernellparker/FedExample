package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"authorization/pkg/db"
	"authorization/pkg/graph/generated"
	"authorization/pkg/graph/models"
	"context"
	"log"
)

func (r *entityResolver) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var users []*models.User
	err := db.Database.Where("username = ?", username).First(&users)
	log.Println(err)
	return users[0], nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
