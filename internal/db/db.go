package database

import (
	"fmt"

	"github.com/SolBaa/finances-backend/internal/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connect() {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	DB = db

	// Opcional: Configurar opciones de GORM
	DB.LogMode(true) // Habilitar registros SQL
}
