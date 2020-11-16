package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"authorization/internal/auth"
	"authorization/internal/middleware"
	"authorization/pkg/graph/generated"
	"authorization/pkg/graph/model"
	"authorization/pkg/graph/models"
	service "authorization/pkg/jwt"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (string, error) {
	var user models.User
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.CreateUser()

	token := service.JWTAuthService().GenerateToken(user.Username, true)
	return token, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input models.Login) (string, error) {
	var user models.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		// 1
		return "", &auth.WrongUsernameOrPasswordError{}
	}
	token := service.JWTAuthService().GenerateToken(user.Username, true)

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input models.RefreshTokenInput) (string, error) {
	token := ""
	//username, err := r.JWTService.ValidateToken(input.Token)
	//if err != nil {
	//	return "", fmt.Errorf("access denied")
	//}
	//token:= r.JWTService.GenerateToken(username,true)
	//if err != nil {
	//	return "", err
	//}
	return token, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var credential models.Login
	ginCtx, ginErr := middleware.GinContextFromContext(ctx)
	if ginErr != nil {
		return []*models.User{}, fmt.Errorf("access denied")
	}
	err := ginCtx.ShouldBind(&credential)
	if err != nil {
		return []*models.User{}, fmt.Errorf("access denied")
	}
	//isUserAuthenticated := controller.loginService.LoginUser(credential.Us, credential.Password)
	//if isUserAuthenticated {
	//	return controller.jWtService.GenerateToken(credential.Email, true)
	//
	//}
	//user := auth.ForContext(ctx)
	//if user == nil {
	//	return []*models.User{}, fmt.Errorf("access denied")
	//}
	return models.GetAllUsers(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
