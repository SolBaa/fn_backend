package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connect(dbString string) *gorm.DB {
	connectionString := dbString

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	} else {
		fmt.Println("Conectado a la base de datos")
	}

	DB = db

	// Opcional: Configurar opciones de GORM
	DB.LogMode(true) // Habilitar registros SQL
	return DB
}
