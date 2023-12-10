package dto

// Fetch
type FetchApiSettingResponse struct {
	ApiSettingId         int    `json:"apiSettingId"`
	ApiId                int    `json:"apiId"`
	RequestMethod        string `json:"requestMethod"`
	Endpoint             string `json:"endpoint"`
	ExecutionIntervalSec int    `json:"executionIntervalSec"`
}

// Add
type AddApiSettingRequest struct {
	ApiId                int    `json:"apiId"`
	RequestMethod        string `json:"requestMethod"`
	Endpoint             string `json:"endpoint"`
	ExecutionIntervalSec int    `json:"executionIntervalSec"`
}

// Update
type UpdateApiSettingRequest struct {
	RequestMethod        string `json:"requestMethod"`
	Endpoint             string `json:"endpoint"`
	ExecutionIntervalSec int    `json:"executionIntervalSec"`
}
