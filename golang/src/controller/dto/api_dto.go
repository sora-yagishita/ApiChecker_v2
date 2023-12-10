package dto

// Fetch
type FetchApiResponse struct {
	ApiId          string `json:"apiId"`
	ApiName        string `json:"apiName"`
	ApiDescription string `json:"apiDescription"`
}

// Add
type AddApiRequest struct {
	ApiName        string `json:"apiName"`
	ApiDescription string `json:"apiDescription"`
}

// Update
type UpdateApiRequest struct {
	ApiName        string `json:"apiName"`
	ApiDescription string `json:"apiDescription"`
}
