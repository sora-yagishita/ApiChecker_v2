package dto

// Fetch
type FetchApiRequest struct {
	ApiID int `json:"api_id"`
}

type FetchApiResponse struct {
	ApiID          int    `json:"api_id"`
	ApiName        string `json:"api_name"`
	ApiDescription string `json:"api_description"`
}

// Add
type AddApiRequest struct {
	ApiName        string `json:"api_name"`
	ApiDescription string `json:"api_description"`
}

// Update
type UpdateApiRequest struct {
	ApiID          int    `json:"api_id"`
	ApiName        string `json:"api_name"`
	ApiDescription string `json:"api_description"`
}

// Delete
type DeleteApiRequest struct {
	ApiID int `json:"api_id"`
}
