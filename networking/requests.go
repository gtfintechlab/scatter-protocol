package networking

import (
	"encoding/json"
	"log"
	"net/http"
)

func PostValidator(request *http.Request, response http.ResponseWriter) {
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusBadRequest)
		log.Fatal("Invalid request method")
		return
	}
}

func GetValidator(request *http.Request, response http.ResponseWriter) {
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusBadRequest)
		log.Fatal("Invalid request method")
		return
	}
}

func SendJson(response http.ResponseWriter, jsonData any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(jsonData)

}
