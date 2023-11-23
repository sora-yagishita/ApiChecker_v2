package dto

// Fetch
type FetchApiRequest struct {
	ApiId string `json:"apiId"`
}

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
	ApiId          string `json:"apiId"`
	ApiName        string `json:"apiName"`
	ApiDescription string `json:"apiDescription"`
}

// Delete
type DeleteApiRequest struct {
	ApiId string `json:"apiId"`
}
