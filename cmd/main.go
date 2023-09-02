package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SolBaa/recipes-backend/cmd/api/handler"
	"github.com/SolBaa/recipes-backend/cmd/api/router"
	database "github.com/SolBaa/recipes-backend/internal/db"
	"github.com/SolBaa/recipes-backend/internal/repository"
	"github.com/SolBaa/recipes-backend/internal/service"
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
	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatal("Error Connecting to DB")
	}
	fmt.Println(db)
	// create router
	// r := chi.NewRouter()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewService(userRepository)
	userHandler := handler.NewUser(userService)
	recipeRepository := repository.NewRecipeRepository(db)
	recipesService := service.NewRecipeService(recipeRepository)
	recipeHandler := handler.NewRecipe(recipesService)
	app := router.NewRouter(db, userHandler, recipeHandler)
	r := app.InitializeRoutes()
	http.ListenAndServe(":8000", r)

}
