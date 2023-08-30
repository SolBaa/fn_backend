package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SolBaa/finances-backend/cmd/api/router"
	database "github.com/SolBaa/finances-backend/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")
	// create connection to DB
	var dbConfig string
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		DBHost     = os.Getenv("DB_HOST")
		DBPort     = os.Getenv("DB_PORT")
		DBUser     = os.Getenv("DB_USER")
		DBPassword = os.Getenv("DB_PASSWORD")
		DBName     = os.Getenv("DB_NAME")
	)

	dbConfig = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DBHost, DBPort, DBUser, DBName, DBPassword)
	db := database.Connect(dbConfig)

	// create router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	router := router.NewRouter(r, db)
	router.MapRoutes()
	http.ListenAndServe(":8000", r)

}
