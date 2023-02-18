package networking

import (
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
