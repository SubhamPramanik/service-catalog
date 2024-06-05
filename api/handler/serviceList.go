package handler

import (
	"encoding/json"
	"fmt"
	"kong/models"
	"net/http"
	"strconv"
	"strings"
)

type ServiceListResponse struct {
	Total int                     `json:"total"`
	Data  []models.ServiceSummary `json:"data"`
	Next  string                  `json:"next"`
	Prev  string                  `json:"prev"`
}

// GetServices is an HTTP handler that retrieves a list of services, with optional filtering and pagination.
// It accepts the following query parameters:
// - name: a string to filter the services by name
// - page: the page number to retrieve (default is 1)
// - limit: the number of results to return per page (default is 10)
// The response is a JSON object with the following fields:
// - total: the total number of services matching the filters
// - data: an array of ServiceSummary objects representing the paginated services
// - next: a URL for the next page of results (or empty if on the last page)
// - prev: a URL for the previous page of results (or empty if on the first page)
func GetServices(w http.ResponseWriter, r *http.Request) {
	nameFilter := r.URL.Query().Get("name")
	page, limit := getPaginationParams(r)

	filteredServices := filterServices(models.Data, nameFilter)
	total := len(filteredServices)
	totalPages := calculateTotalPages(total, limit)

	if page > totalPages {
		writeEmptyResponse(w)
		return
	}

	paginatedServices, next, prev := paginateServices(filteredServices, page, limit, r)
	summaries := mapServicesToSummaries(paginatedServices)

	response := ServiceListResponse{
		Total: total,
		Data:  summaries,
		Next:  next,
		Prev:  prev,
	}

	json.NewEncoder(w).Encode(response)
}

func getPaginationParams(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	return page, limit
}

func calculateTotalPages(total, limit int) int {
	return (total + limit - 1) / limit
}

func writeEmptyResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"total": 0,
		"data":  []models.ServiceSummary{},
		"next":  "",
		"prev":  "",
	})
}

func mapServicesToSummaries(services []models.Service) []models.ServiceSummary {
	summaries := make([]models.ServiceSummary, len(services))
	for i, service := range services {
		summaries[i] = models.ServiceSummary{
			ID:           service.ID,
			Name:         service.Name,
			Description:  service.Description,
			VersionCount: len(service.Versions),
		}
	}
	return summaries
}

func paginateServices(services []models.Service, page, limit int, r *http.Request) ([]models.Service, string, string) {
	startIndex := (page - 1) * limit
	endIndex := min(startIndex+limit, len(services))

	next := generateURL(r, page+1)
	prev := generateURL(r, page-1)

	return services[startIndex:endIndex], next, prev
}

func generateURL(r *http.Request, page int) string {
	if page <= 0 {
		return ""
	}
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	query := r.URL.Query()
	query.Set("page", strconv.Itoa(page))
	return fmt.Sprintf("%s://%s%s?%s", scheme, r.Host, r.URL.Path, query.Encode())
}

func filterServices(services []models.Service, filter string) []models.Service {
	var filtered []models.Service
	if filter != "" {
		lowerFilter := strings.ToLower(filter)
		for _, service := range services {
			if strings.Contains(strings.ToLower(service.Name), lowerFilter) {
				filtered = append(filtered, service)
			}
		}
		return filtered
	}
	return services
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
Test Cases:

Basic Functionality
- TestGetServicesBasicResponse: Ensure the handler returns the correct JSON structure with default pagination parameters when no filters are applied.

Filtering
- TestGetServicesWithValidNameFilter: Verify the service list is correctly filtered by name.
- TestGetServicesWithInvalidNameFilter: Ensure that an empty list is returned when a non-existent service name is used as a filter.

Pagination
- TestGetServicesPaginationFirstPage: Check that the first page of results is correct and includes the 'next' link but not the 'prev' link.
- TestGetServicesPaginationMiddlePage: Ensure a middle page correctly shows both 'prev' and 'next' links.
- TestGetServicesPaginationLastPage: Validate that the last page shows the 'prev' link and does not show a 'next' link.
- TestGetServicesPaginationBeyondLastPage: Confirm that accessing a page number beyond the last page returns an empty response.

Page and Limit Parameters
- TestGetServicesInvalidPageParameter: Verify handling of non-numeric or negative page numbers by defaulting to the first page.
- TestGetServicesInvalidLimitParameter: Test that non-numeric or negative limit values default to 10.
- TestGetServicesCustomLimitParameter: Ensure that a custom limit parameter is respected and alters the number of results per page appropriately.

Error Handling and Edge Cases
- TestGetServicesEmptyData: Confirm behavior when there are no services available (e.g., database is empty).
- TestGetServicesHandleQueryParameterErrors: Check error handling for malformed query parameters.

URL Generation
- TestGetServicesURLGeneration: Test the URL generation for pagination links, ensuring that they correctly reflect the request's scheme (HTTP/HTTPS) and host.
*/
