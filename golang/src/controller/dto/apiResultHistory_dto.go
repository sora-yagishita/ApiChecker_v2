package dto

import (
	"time"
)

// Fetch
type FetchApiResultHistoryResult struct {
	ApiId              int       `json:"apiId"`
	ApiResultId        int       `json:"apiResultId"`
	RequestEndpoint    string    `json:"requestEndpoint"`
	RequestParamString string    `json:"requestParamString"`
	RequestDateTime    time.Time `json:"requestDateTime"`
	ResponseStatusCode string    `json:"responseStatusCode"`
	ResponseData       string    `json:"responseData"`
	ResponseDateTime   time.Time `json:"responseDateTime"`
}

// Add
type AddApiResultHistoryRequest struct {
	ApiId              int       `json:"apiId"`
	RequestEndpoint    string    `json:"requestEndpoint"`
	RequestParamString string    `json:"requestParamString"`
	RequestDateTime    time.Time `json:"requestDateTime"`
	ResponseStatusCode string    `json:"responseStatusCode"`
	ResponseData       string    `json:"responseData"`
	ResponseDateTime   time.Time `json:"responseDateTime"`
}

// Update
type UpdateApiResultHistoryRequest struct {
	RequestEndpoint    string    `json:"requestEndpoint"`
	RequestParamString string    `json:"requestParamString"`
	RequestDateTime    time.Time `json:"requestDateTime"`
	ResponseStatusCode string    `json:"responseStatusCode"`
	ResponseData       string    `json:"responseData"`
	ResponseDateTime   time.Time `json:"responseDateTime"`
}
