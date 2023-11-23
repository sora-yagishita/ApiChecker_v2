package dto

// Fetch
type FetchApiParamSettingRequest struct {
	ApiSettingID int `json:"apiSettingId"`
}

type FetchApiParamSettingResponse struct {
	ApiSettingID      int    `json:"apiSettingId"`
	ApiParamSettingNo int    `json:"apiParamSettingNo"`
	ApiParamKey       string `json:"apiParamKey"`
	ApiParamValue     string `json:"apiParamValue"`
}

// Add
type AddApiParamSettingRequest struct {
	ApiSettingID  int    `json:"apiSettingId"`
	ApiParamKey   string `json:"apiParamKey"`
	ApiParamValue string `json:"apiParamValue"`
}

// Update
type UpdateApiParamSettingRequest struct {
	ApiSettingID  int    `json:"apiSettingId"`
	ApiParamKey   string `json:"apiParamKey"`
	ApiParamValue string `json:"apiParamValue"`
}

// Delete
type DeleteApiParamSettingRequest struct {
	ApiSettingID int `json:"apiSettingId"`
}
