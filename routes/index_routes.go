package routes

import (
	"github.com/gorilla/mux"
	"github.com/juanmontilva/api_golang/controllers"
)

func InitRouter() *mux.Router {
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", controllers.GetInitRoute).Methods("GET")
	api.HandleFunc("", controllers.GetInitRoute).Methods("GET")

	return rutas
}
