package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api-monito/graph/generated"
	"api-monito/graph/model"
	"context"
	"log"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUser) (*model.User, error) {
	userData, errCreateUser := r.userService.CreateAccount(input.EmailAddress, input.Password)

	if errCreateUser != nil{
		log.Println(errCreateUser.Error())
		return &model.User{}, errCreateUser
	}

	return &model.User{
		ID: userData.ID,
		EmailAddress: userData.EmailAddress,
		HashedUserID: userData.HashedUserId,
		PasswordHash: userData.PasswordHash,
		Projects: userData.Projects,
		APIKey: userData.ApiKey,
		TimeZone: userData.TimeZone,
		NotificationTurnedOn: userData.NotificationTurnedOn,
	},nil 

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{
	 *Resolver

}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type queryResolver struct{ *Resolver}
