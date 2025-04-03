package db

import (
	"log" // Paquete para registrar mensajes en la consola
	"os"  // Paquete para trabajar con variables de entorno

	"gorm.io/driver/postgres" // Driver de PostgreSQL para GORM
	"gorm.io/gorm"            // ORM (Object Relational Mapper) para interactuar con la base de datos
	"gorm.io/gorm/logger"     // Logger para GORM
)

// ConectarPostgres establece una conexión con la base de datos PostgreSQL
// y devuelve una instancia de *gorm.DB o un error en caso de fallo.
func ConectarPostgres() (*gorm.DB, error) {
	// Obtiene la cadena de conexión desde la variable de entorno POSTGRES_CONNECTION_STRING
	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connectionString == "" {
		// Si la variable de entorno no está configurada, registra un mensaje y devuelve un error
		log.Println("La variable de entorno POSTGRES_CONNECTION_STRING no está configurada")
		return nil, gorm.ErrInvalidDB
	}

	// Intenta abrir una conexión con la base de datos usando GORM y el driver de PostgreSQL
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		// Configura el logger de GORM para mostrar mensajes de nivel "Info"
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		// Si ocurre un error al conectar, devuelve el error
		return nil, err
	}

	// Si la conexión es exitosa, registra un mensaje en la consola
	log.Println("Conectado a la base de datos")
	// Devuelve la instancia de la base de datos y nil como error
	return db, nil
}

// ESTO FUNCIONA PARA CONECTAR EN LA BASE DE DATOS DE MANERA SEGURA, OJO AGREGAR CODIGO EXPORT A LA TERMINAL Y ASI QUEDA TODO OK MUY EXCELENTE

// export POSTGRES_CONNECTION_STRING="host=localhost user=postgres password=1234 dbname=api_golang port=5433 sslmode=disable TimeZone=America/Caracas"

// var CONNECTION_STRING = "host=localhost user=postgres password=1234 dbname=api_golang port=5433 sslmode=disable TimeZone=America/Caracas"
// var DB *gorm.DB

// ESTE ES EL CODGIO ANTERIOR, SE MEJORO MUCHO LA METODOLOGIA DE SEGURIDAD, ES EXCELENTE

// "gorm.io/driver/postgres"
// "gorm.io/gorm"
//

//ESTE CODIGO ES EL ANTERIOR.... SE MEJORO MUCHO EL PERFORMANCE DE GORM PARA CONECTAR A LA BASE DE DATOS

// var CONNECTION_STRING = "host=localhost user=postgres password=1234 dbname=api_golang port=5433 sslmode=disable TimeZone=America/Caracas"
// var DB *gorm.DB

// func ConectarPostgres() {
// 	var error error
// 	DB, error = gorm.Open(postgres.Open(CONNECTION_STRING), &gorm.Config{})
// 	if error != nil {
// 		log.Fatal("ERROR DE CONEXION A LA BASE DE DATOS: ", error)
// 	} else {
// 		log.Println("conectado a la bd")
// 	}
// }
