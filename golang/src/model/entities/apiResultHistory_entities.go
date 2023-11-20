package entities

import (
	"time"
)

type ApiResultHistory struct {
	ApiResultID        int       `json:"api_result_id"`
	ApiID              int       `json:"api_id"`
	RequestEndpoint    string    `json:"request_endpoint"`
	RequestParamString string    `json:"request_param_string"`
	RequestDateTime    time.Time `json:"request_date_time"`
	ResponseStatusCode string    `json:"response_status_code"`
	ResponseData       string    `json:"response_data"`
}
