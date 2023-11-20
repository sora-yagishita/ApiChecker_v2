package dto

// Fetch
type FetchApiHeaderSettingRequest struct {
	ApiSettingID int `json:"api_setting_id"`
}

type FetchApiHeaderSettingResponse struct {
	ApiSettingID       int    `json:"api_setting_id"`
	ApiHeaderSettingNo int    `json:"api_header_setting_no"`
	ApiHeaderKey       string `json:"api_header_key"`
	ApiHeaderValue     string `json:"api_header_value"`
}

// Add
type AddApiHeaderSettingRequest struct {
	ApiSettingID   int    `json:"api_setting_id"`
	ApiHeaderKey   string `json:"api_header_key"`
	ApiHeaderValue string `json:"api_header_value"`
}

// Update
type UpdateApiHeaderSettingRequest struct {
	ApiSettingID   int    `json:"api_setting_id"`
	ApiHeaderKey   string `json:"api_header_key"`
	ApiHeaderValue string `json:"api_header_value"`
}

// Delete
type DeleteApiHeaderSettingRequest struct {
	ApiSettingID int `json:"api_setting_id"`
}
