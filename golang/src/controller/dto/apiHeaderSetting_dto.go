package dto

// Fetch
type FetchApiHeaderSettingRequest struct {
	ApiSettingId int `json:"apiSettingId"`
}

type FetchApiHeaderSettingResponse struct {
	ApiSettingId       int    `json:"apiSettingId"`
	ApiHeaderSettingNo int    `json:"apiHeaderSettingNo"`
	ApiHeaderKey       string `json:"apiHeaderKey"`
	ApiHeaderValue     string `json:"apiHeaderValue"`
}

// Add
type AddApiHeaderSettingRequest struct {
	ApiSettingId       int    `json:"apiSettingId"`
	ApiHeaderSettingNo int    `json:"apiHeaderSettingNo"`
	ApiHeaderKey       string `json:"apiHeaderKey"`
	ApiHeaderValue     string `json:"apiHeaderValue"`
}

// Update
type UpdateApiHeaderSettingRequest struct {
	ApiHeaderKey   string `json:"apiHeaderKey"`
	ApiHeaderValue string `json:"apiHeaderValue"`
}

// Delete
type DeleteApiHeaderSettingRequest struct {
	ApiSettingId int `json:"apiSettingId"`
}
