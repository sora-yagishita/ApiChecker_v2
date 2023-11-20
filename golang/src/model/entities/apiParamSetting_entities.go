package entities

type ApiParamSetting struct {
	ApiSettingID      int    `json:"api_setting_id"`
	ApiParamSettingNo int    `json:"api_param_setting_no"`
	ApiParamKey       string `json:"api_param_key"`
	ApiParamValue     string `json:"api_param_value"`
}
