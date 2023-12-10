package entities

type ApiHeaderSetting struct {
	ApiSettingId       int    `json:"api_setting_id"`
	ApiHeaderSettingNo int    `json:"api_header_setting_no"`
	ApiHeaderKey       string `json:"api_header_key"`
	ApiHeaderValue     string `json:"api_header_value"`
}
