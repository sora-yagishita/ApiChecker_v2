package dto

type ApiSetting struct {
	ApiSettingID         int    `json:"apiSetting_id"`
	ApiID                int    `json:"apiId"`
	RequestMethod        string `json:"requestMethod"`
	Endpoint             string `json:"endpoint"`
	ExecutionIntervalSec int    `json:"executionIntervalSec"`
}
