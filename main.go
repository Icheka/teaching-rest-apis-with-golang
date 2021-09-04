package main

import (
	"log"
	"net/http"
	"teaching-rest-apis-with-golang/database"
	"teaching-rest-apis-with-golang/routes"

	"github.com/gorilla/mux"
)

func initialiseDatabase() {
	database.PopulateDB()
}

func main() {
	initialiseDatabase()
	router := mux.NewRouter()

	router.HandleFunc("/", routes.Welcome).Methods("GET")
	// CRUD ENDPOINTS
	// CREATE new plate
	router.HandleFunc("/plates", routes.CreatePlate).Methods("POST")
	// GET all plates
	router.HandleFunc("/plates", routes.GetAllPlates).Methods("GET")
	// UPDATE plate by ID
	router.HandleFunc("/plates/{id}", routes.UpdatePlateById).Methods("PUT")
	// DELETE plate by ID
	router.HandleFunc("/plates/{id}", routes.DeletePlate).Methods("DELETE")

	log.Fatal((http.ListenAndServe(":8000", router)))
}
