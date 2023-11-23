package dto

// Fetch
type FetchApiHeaderSettingRequest struct {
	ApiSettingID int `json:"apiSettingId"`
}

type FetchApiHeaderSettingResponse struct {
	ApiSettingID       int    `json:"apiSettingId"`
	ApiHeaderSettingNo int    `json:"apiHeaderSettingNo"`
	ApiHeaderKey       string `json:"apiHeaderKey"`
	ApiHeaderValue     string `json:"apiHeaderValue"`
}

// Add
type AddApiHeaderSettingRequest struct {
	ApiSettingID   int    `json:"apiSettingId"`
	ApiHeaderKey   string `json:"apiHeaderKey"`
	ApiHeaderValue string `json:"apiHeaderValue"`
}

// Update
type UpdateApiHeaderSettingRequest struct {
	ApiSettingID   int    `json:"apiSettingId"`
	ApiHeaderKey   string `json:"apiHeaderKey"`
	ApiHeaderValue string `json:"apiHeaderValue"`
}

// Delete
type DeleteApiHeaderSettingRequest struct {
	ApiSettingID int `json:"apiSettingId"`
}
