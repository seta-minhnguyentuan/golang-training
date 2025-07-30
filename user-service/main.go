package main

import (
	"log"
	"net/http"
	"user-service/config"
	"user-service/graph"
	"user-service/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env")
	}
}

func main() {
	port := config.Getenv("PORT", "8080")
	db := config.ConnectDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
