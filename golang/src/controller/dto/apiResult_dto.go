package dto

import (
	"time"
)

type AddApiResultRequest struct {
	ApiName string `json:"apiName"`
	ApiStatus string `json:"apiStatus"`
	ApiDateTime time.Time `json:"apiDateTime"`
}