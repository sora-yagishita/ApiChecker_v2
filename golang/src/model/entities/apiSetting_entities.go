package entities

type ApiSetting struct {
	ApiSettingId         int    `json:"api_setting_id"`
	ApiId                int    `json:"api_id"`
	RequestMethod        string `json:"request_method"`
	Endpoint             string `json:"endpoint"`
	ExecutionIntervalSec int    `json:"execution_interval_sec"`
}
