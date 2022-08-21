package userService

import (
	"api-monito/models/User"
	"api-monito/repository"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)


type IUserService interface{

}

type UserService struct {
	userRepo repository.IUserRepository
}

func InitUserService(userRepo repository.IUserRepository) *UserService{
	return &UserService{
		userRepo: userRepo,
	}
}

func(userService UserService) CreateAccount(emailAddress, password string) (User.User, error ) {
	doesUserExist, user, errorGetUser :=userService.userRepo.GetUserByEmail(emailAddress)

	if errorGetUser != nil{
		log.Println(errorGetUser.Error())
		return user, errorGetUser
	}

	if doesUserExist == true {
		return user,  errors.New("user already exists")
	}


	bytes, errorHashPass := bcrypt.GenerateFromPassword([]byte(password), 14)

	if errorHashPass != nil{
		log.Println(errorHashPass.Error())
		return User.User{}, errorHashPass
	}

	passwordHash := string(bytes)

	hasCreated, userObject, errCreateUser :=userService.userRepo.Create(emailAddress, passwordHash)

	if errCreateUser != nil{
		log.Println(errCreateUser.Error())
		return User.User{}, errCreateUser
	}

	if hasCreated == false {
		return User.User{}, errors.New("unable to create user")
	}

	return userObject, nil 



}


