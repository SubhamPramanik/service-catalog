package handler

import (
	"encoding/json"
	"kong/models"
	"net/http"

	"github.com/gorilla/mux"
)

// ErrorResponse represents the JSON structure for an error message.
type ErrorResponse struct {
	Error string `json:"error"`
}

// GetServiceByID retrieves a service by its ID and returns it as JSON. If not found, it returns a 404 with a custom error message.
func GetServiceByID(w http.ResponseWriter, r *http.Request) {
	serviceID := mux.Vars(r)["id"]
	service, found := findServiceByID(serviceID)
	if !found {
		writeJSONError(w, "Service not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(service)
}

// GetVersionByID retrieves a specific version of a service by service ID and version ID, returning it as JSON.
// If the service or the version is not found, it returns a 404 with a custom error message.
func GetVersionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service, found := findServiceByID(params["id"])
	if !found {
		writeJSONError(w, "Service not found", http.StatusNotFound)
		return
	}

	version, found := findVersionByID(service, params["versionId"])
	if !found {
		writeJSONError(w, "Version not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(version)
}

func findServiceByID(id string) (*models.Service, bool) {
	for _, service := range models.Data {
		if service.ID == id {
			return &service, true
		}
	}
	return nil, false
}

func findVersionByID(service *models.Service, versionID string) (*models.Version, bool) {
	for _, version := range service.Versions {
		if version.ID == versionID {
			return &version, true
		}
	}
	return nil, false
}

func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

/**
Test Cases:

GetServiceByID Functionality
Basic Functionality
- TestGetServiceByIDFound: Ensure that a valid service ID returns the correct service as JSON.
- TestGetServiceByIDNotFound: Test that an invalid service ID returns a 404 error with the correct error message.
Edge Cases and Error Handling
- TestGetServiceByIDEmptyID: Verify that requesting with an empty ID parameter correctly handles the error.

GetVersionByID Functionality
Basic Functionality
- TestGetVersionByIDFound: Ensure that valid service and version IDs return the correct version as JSON.
- TestGetVersionByIDServiceNotFound: Test that a non-existent service ID returns a 404 error with the appropriate error message.
- TestGetVersionByIDVersionNotFound: Check that a non-existent version ID for an existing service returns a 404 error with the appropriate error message.
Edge Cases and Error Handling
- TestGetVersionByIDEmptyServiceID: Verify error handling when the service ID is empty.
- TestGetVersionByIDEmptyVersionID: Check error handling when the version ID is empty.

Common Error Handling
- TestErrorResponseFormat: Validate that the error response JSON structure matches the expected format (including content type and status codes).
*/
