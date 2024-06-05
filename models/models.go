package models

// TODO: There should be a createdAt and updatedAt field
// which is normally present in databases. We should allow
// filtering and sorting based on the timestamps as well.
type Service struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Versions    []Version `json:"versions"`
}

type Version struct {
	ID         string `json:"id"`
	Number     string `json:"number"`
	CreatedBy  string `json:"createdBy"`
	CommitHash string `json:"commitHash"`
}

type ServiceSummary struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	VersionCount int    `json:"versionCount"`
}
