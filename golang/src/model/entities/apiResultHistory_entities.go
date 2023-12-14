package entities

import (
	"time"
)

type ApiResultHistory struct {
	ApiResultId        int       `json:"api_result_id"`
	ApiId              int       `json:"api_id"`
	RequestEndpoint    string    `json:"request_endpoint"`
	RequestParamString string    `json:"request_param_string"`
	RequestDateTime    time.Time `json:"request_date_time"`
	ResponseStatusCode string    `json:"response_status_code"`
	ResponseData       string    `json:"response_data"`
	ResponseDateTime   time.Time `json:"response_date_time"`
}
