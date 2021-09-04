package routes

import "net/http"

func SetJsonHeader(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
}
