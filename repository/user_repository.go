package repository

import (
	"api-monito/models/User"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	userTable *gorm.DB

}

func NewUserRepository(db (*gorm.DB)) *UserRepository {
	return &UserRepository{
		userTable: db.Table("users"),
	}
}

func (userRepo *UserRepository) Create(emailAddress, password string) (User.User, error ){
	newUserId:= uuid.NewString()
	newAPIKey :=  uuid.NewString()
	newUser := User.User{
		ID: newUserId ,
		EmailAddress: emailAddress,
		HashedUserId: newUserId,
		PasswordHash: password,
		Projects: make([]string, 0),
		ApiKey: newAPIKey,
		TimeZone: " ",
		NotificationTurnedOn: false,
	}

	result := userRepo.userTable.Create(newUser)

	if result.Error != nil{
		log.Println(result.Error.Error())
		return User.User{}, result.Error
	}

	return newUser, nil 
}