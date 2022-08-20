package main

import (
	"api-monito/graph"
	"api-monito/graph/generated"
	"api-monito/models/User"
	"api-monito/utils/dbUtil"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, errStartDB := dbUtil.SetUpDB()

	db.AutoMigrate(&User.User{})

	if errStartDB != nil{
		log.Println(errStartDB.Error())
		return 
	}

	

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
