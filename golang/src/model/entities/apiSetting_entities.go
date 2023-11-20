package entities

type ApiSetting struct {
	ApiSettingID         int    `json:"api_setting_id"`
	ApiID                int    `json:"api_id"`
	RequestMethod        string `json:"request_method"`
	Endpoint             string `json:"endpoint"`
	ExecutionIntervalSec int    `json:"execution_interval_sec"`
}
