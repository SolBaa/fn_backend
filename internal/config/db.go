package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func GetDBConfig() string {
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
	return dbConfig
}
