package main

import (
	"log"
	"net/http"

	"github.com/juanmontilva/api_golang/db"

	//"github.com/gorilla/mux"
	"github.com/juanmontilva/api_golang/routes"
)

func main() {
	db.ConectarPostgres()
	rutas := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", rutas))

}
