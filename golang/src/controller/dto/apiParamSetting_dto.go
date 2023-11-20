package dto

// Fetch
type FetchApiParamSettingRequest struct {
	ApiSettingID int `json:"api_setting_id"`
}

type FetchApiParamSettingResponse struct {
	ApiSettingID      int    `json:"api_setting_id"`
	ApiParamSettingNo int    `json:"api_param_setting_no"`
	ApiParamKey       string `json:"api_param_key"`
	ApiParamValue     string `json:"api_param_value"`
}

// Add
type AddApiParamSettingRequest struct {
	ApiSettingID  int    `json:"api_setting_id"`
	ApiParamKey   string `json:"api_param_key"`
	ApiParamValue string `json:"api_param_value"`
}

// Update
type UpdateApiParamSettingRequest struct {
	ApiSettingID  int    `json:"api_setting_id"`
	ApiParamKey   string `json:"api_param_key"`
	ApiParamValue string `json:"api_param_value"`
}

// Delete
type DeleteApiParamSettingRequest struct {
	ApiSettingID int `json:"api_setting_id"`
}
