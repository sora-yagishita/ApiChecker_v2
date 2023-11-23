package dto

import (
	"time"
)

// Fetch
type FetchApiResultHistoryResult struct {
	ApiResultID        int       `json:"apiResultId"`
	ApiID              int       `json:"apiId"`
	RequestEndpoint    string    `json:"requestEndpoint"`
	RequestParamString string    `json:"requestParamString"`
	RequestDateTime    time.Time `json:"requestDateTime"`
	ResponseStatusCode string    `json:"responseStatusCode"`
	ResponseData       string    `json:"responseData"`
}
