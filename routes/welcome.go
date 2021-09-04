package routes

import (
	"encoding/json"
	"net/http"
)

func Welcome(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode("Welcome to Koala, the one-stop online marketplace for your household needs")
}
