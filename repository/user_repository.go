package repository

import (
	"api-monito/models/User"
	"api-monito/utils/stringUtility"
	"log"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface{
	Create(emailAddress , password string) (bool, User.User, error)
	GetOne(userId string)(User.User, error)
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

func (userRepo *UserRepository) Create(emailAddress, password string) (User.User, error ){
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

	result := userRepo.userTable.Create(newUser)

	if result.Error != nil{
		log.Println(result.Error.Error())
		return User.User{}, result.Error
	}

	return newUser, nil 
}

func(userRepo *UserRepository) GetOne(userId string)(bool,  User.User, error ){

	var user User.User
	result := userRepo.userTable.First(&user, "id = ?",  userId)

	if result.Error != nil{
		log.Println(result.Error.Error())
		return false,  User.User{}, result.Error
	}

	return result.RowsAffected > 0, user, nil 
}

func(userRepo *UserRepository)GetUserByEmail(emailAddress string)(bool, User.User, error ){
	var user User.User

	result := userRepo.userTable.Where("emailAddress = ?", emailAddress).First(&user)

	if result.Error != nil{
		log.Println(result.Error.Error())
		return false, user, result.Error
	}

	return result.RowsAffected > 0, user, nil 
}