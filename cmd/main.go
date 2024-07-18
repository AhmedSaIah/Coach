package main

import (
    "log"
    "net/http"
    "TaskManager/handlers"
    "github.com/gorilla/mux"
    "fmt"
    "os"

    "TaskManager/models"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
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

func main() {
    initDB()
    migrate()

    r := mux.NewRouter()
    taskHandler := handlers.TaskHandler{DB: db}

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, "<h1>Welcome to Task-Management System API</h1>")
    }).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
    r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
    r.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
    r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
