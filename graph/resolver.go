package graph

import (
	"api-monito/graph/model"
	"api-monito/models/User"
	"api-monito/services/userService"
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	userService userService.IUserService
}

func NewResolver(userService userService.IUserService) *Resolver {
	return &Resolver{
		userService: userService,
	}
}

func(resolver *Resolver) CreateUser(ctx context.Context, input model.CreateUser)(User.User, error){
	return resolver.userService.CreateAccount(input.EmailAddress, input.Password)
}



