package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"teaching-rest-apis-with-golang/database"

	"github.com/gorilla/mux"
)

type TNewPlate struct {
	Color      string  `json:"color"`
	Dimensions string  `json:"dimension"`
	Price      float64 `json:"price"`
}

// CREATE new plate
func CreatePlate(res http.ResponseWriter, req *http.Request) {
	SetJsonHeader(res)

	var newPlate TNewPlate
	_ = json.NewDecoder(req.Body).Decode(&newPlate)

	var plate database.Plate = database.Plate{
		Id:         strconv.Itoa(len(database.Plates) + 1),
		Color:      newPlate.Color,
		Dimensions: newPlate.Dimensions,
		Price:      newPlate.Price,
	}

	database.Plates = append(database.Plates, plate)
	json.NewEncoder(res).Encode(database.Plates)
}

// GET all plates
func GetAllPlates(res http.ResponseWriter, req *http.Request) {
	SetJsonHeader(res)
	json.NewEncoder(res).Encode(database.Plates)
}

// UPDATE plate by Id field
func UpdatePlateById(res http.ResponseWriter, req *http.Request) {
	SetJsonHeader(res)
	params := mux.Vars(req)

	for index, plate := range database.Plates {
		if plate.Id == params["id"] {
			// remove the matching plate from the database, and...
			database.Plates = append(database.Plates[:index], database.Plates[index+1:]...)
			// append the updated plate
			var newPlate database.Plate
			_ = json.NewDecoder(req.Body).Decode(&newPlate)
			database.Plates = append(database.Plates, newPlate)
			json.NewEncoder(res).Encode(database.Plates)
			return
		}
	}

	json.NewEncoder(res).Encode("Not found")
}

// DELETE plate by Id
func DeletePlate(res http.ResponseWriter, req *http.Request) {
	SetJsonHeader(res)

	params := mux.Vars(req)

	for index, plate := range database.Plates {
		if plate.Id == params["id"] {
			database.Plates = append(database.Plates[:index], database.Plates[index+1:]...)
			json.NewEncoder(res).Encode(database.Plates)
			return
		}
	}
}
