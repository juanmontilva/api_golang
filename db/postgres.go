package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var CONNECTION_STRING = "host=localhost user=postgres password=1234 dbname=api_golang port=5433 sslmode=disable TimeZone=America/Caracas"
var DB *gorm.DB

func ConectarPostgres() {
	var error error
	DB, error = gorm.Open(postgres.Open(CONNECTION_STRING), &gorm.Config{})
	if error != nil {
		log.Fatal("ERROR DE CONEXION A LA BASE DE DATOS: ", error)
	} else {
		log.Println("conectado a la bd")
	}
}
