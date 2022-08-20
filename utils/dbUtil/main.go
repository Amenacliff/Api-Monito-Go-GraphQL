package dbUtil

import (
	"api-monito/constants"
	"api-monito/utils/envUtil"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDB()(*gorm.DB, error ){
	postgresUrl, errorGetEnv := envUtil.GetENVValue(constants.POSTGRES_ENV_KEY)

	if errorGetEnv != nil{
		log.Println(errorGetEnv.Error())
		return  &gorm.DB{}, errorGetEnv
	}

	log.Println(postgresUrl)

	db, err := gorm.Open(postgres.Open(postgresUrl), &gorm.Config{})

	if err != nil{
		log.Println(err.Error())
		return &gorm.DB{}, err
	}

	return  db, nil 
}