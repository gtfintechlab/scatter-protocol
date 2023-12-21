package networking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func SendJson(response http.ResponseWriter, jsonData map[string]interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(jsonData)

}

// MakePostRequestToServer makes a POST request to the provided http.Server.
func MakePostRequestToServer(server *http.Server, path string, data []byte) ([]byte, error) {
	// Construct the URL using the Server's Address and the given path
	url := fmt.Sprintf("http://%s%s", server.Addr, path)

	// Create a new request with POST method, URL, and data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %v", err)
	}

	// Set headers if needed (e.g., content-type)
	req.Header.Set("Content-Type", "application/json") // Modify this if your data format is different

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform POST request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, response body: %s", resp.StatusCode, body)
	}

	return body, nil
}
