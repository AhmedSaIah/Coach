package main

import (
	"TaskManager/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

    // "TaskManager/graph"
	// "TaskManager/graph/resolver"

    "net/http"
    // "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
)

var db *gorm.DB

func initDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	databaseUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDb := os.Getenv("POSTGRES_DB")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	if databaseUser == "" || postgresPassword == "" || postgresDb == "" {
		log.Fatal("Missing required environment variables")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, databaseUser, postgresPassword, postgresDb, port)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected")
}

func migrate() {
	err := db.AutoMigrate(&models.Task{})
	if err != nil {
		return
	}
}

// const defaultPort = "8080"

func main() {
    initDB()
    migrate()

    // srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{DB: db}}))

    // http.Handle("/graphql", srv)
    http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

    log.Println("connect to http://localhost:8080/playground for GraphQL playground")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
