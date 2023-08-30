package main

import (
	"fmt"

	database "github.com/SolBaa/finances-backend/internal/db"
)

func main() {
	fmt.Println("Hello, World!")
	// create connection to DB
	database.Connect()
	// defer close connection to DB
	// create new user repository
	// create new user service
}
