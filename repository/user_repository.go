package repository

import (
	"api-monito/models/User"
	"api-monito/utils/stringUtility"
	"context"
	"log"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface{
	Create(emailAddress , password string) (bool, User.User, error)
	GetOne(userId string)(bool, User.User, error)
	GetUserByEmail(emailAddress string)(bool, User.User, error )
}

type UserRepository struct {
	userTable *gorm.DB

}

func NewUserRepository(db (*gorm.DB)) *UserRepository {
	return &UserRepository{
		userTable: db.Table("users"),
	}
}

func (userRepo *UserRepository) Create(emailAddress, password string) (bool, User.User, error ){
	newUserId:= uuid.NewString()
	newAPIKey :=  uuid.NewString()
	newUser := User.User{
		ID: newUserId ,
		EmailAddress: strings.ToLower(stringUtility.TrimALLSpace(emailAddress)),
		HashedUserId: newUserId,
		PasswordHash: password,
		Projects: make([]string, 0),
		ApiKey: newAPIKey,
		TimeZone: " ",
		NotificationTurnedOn: false,
	}


	result := userRepo.userTable.WithContext(context.Background()).Create(&newUser)

		log.Println(result)


	if result.Error != nil{
		log.Println(result.Error.Error())
		return false,  User.User{}, result.Error
	}

	return result.RowsAffected > 0,  newUser, nil 
}

func(userRepo *UserRepository) GetOne(userId string)(bool,  User.User, error ){

	var user User.User
	result := userRepo.userTable.WithContext(context.Background()).First(&user, "id = ?",  userId)

	if result.Error != nil{
		log.Println(result.Error.Error())
		return false,  User.User{}, result.Error
	}

	return result.RowsAffected > 0, user, nil 
}

func(userRepo *UserRepository)GetUserByEmail(emailAddress string)(bool, User.User, error ){
	var user User.User

	result := userRepo.userTable.WithContext(context.Background()).Where(User.User{EmailAddress: emailAddress}).First(&user)

	if result.Error != nil{
		if(result.Error.Error()  == "record not found"){
			return false, user, nil 
		}
		return false, user, result.Error
		}
		
	return result.RowsAffected > 0, user, nil 
}