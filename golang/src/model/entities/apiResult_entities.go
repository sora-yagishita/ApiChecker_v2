package entities

import (
	"time"
)

type ApiResult struct {
	ApiName string `json:"apiName"`
	ApiStatus string `json:"apiStatus"`
	ApiDateTime time.Time `json:"apiDateTime"`
}