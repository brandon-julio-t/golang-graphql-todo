package main

import (
	"github.com/brandon-julio-t/golang-graphql-todo/app/repositories"
	"github.com/brandon-julio-t/golang-graphql-todo/app/services"
	"github.com/brandon-julio-t/golang-graphql-todo/graph/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brandon-julio-t/golang-graphql-todo/graph"
	"github.com/brandon-julio-t/golang-graphql-todo/graph/generated"
)

const defaultPort = "8080"

func main() {
	port, db := setup()
	srv := setupServer(db)
	runServer(srv, port)
}

func setup() (string, *gorm.DB) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	databaseUrl := os.Getenv("DATABASE_URL")

	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open(
		postgres.Open(databaseUrl),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		panic(err)
	}

	if err := db.Debug().AutoMigrate(&model.Todo{}); err != nil {
		panic(err)
	}

	return port, db
}

func setupServer(db *gorm.DB) *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					TodoService: &services.TodoService{
						Repository: &repositories.TodoRepository{
							DB: db,
						},
					},
				},
			},
		),
	)
}

func runServer(srv *handler.Server, port string) {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/playground"))
	http.Handle("/graphql", srv)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
