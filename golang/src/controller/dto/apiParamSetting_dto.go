package dto

// Fetch
type FetchApiParamSettingRequest struct {
	ApiSettingId int `json:"apiSettingId"`
}

type FetchApiParamSettingResponse struct {
	ApiSettingId      int    `json:"apiSettingId"`
	ApiParamSettingNo int    `json:"apiParamSettingNo"`
	ApiParamKey       string `json:"apiParamKey"`
	ApiParamValue     string `json:"apiParamValue"`
}

// Add
type AddApiParamSettingRequest struct {
	ApiSettingId      int    `json:"apiSettingId"`
	ApiParamSettingNo int    `json:"apiParamSettingNo"`
	ApiParamKey       string `json:"apiParamKey"`
	ApiParamValue     string `json:"apiParamValue"`
}

// Update
type UpdateApiParamSettingRequest struct {
	ApiParamKey   string `json:"apiParamKey"`
	ApiParamValue string `json:"apiParamValue"`
}

// Delete
type DeleteApiParamSettingRequest struct {
	ApiSettingId int `json:"apiSettingId"`
}
