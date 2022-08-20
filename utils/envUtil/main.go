package envUtil

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func GetENVValue(key string)(string, error ){
	err := godotenv.Load()

	if err != nil{
		log.Println(err.Error())
		return "", err
	}

	value := os.Getenv(key)

	return value, nil 

}