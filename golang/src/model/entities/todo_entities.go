package entities

import (
	"time"
)

type Todo struct {
	ApiName     string    `json:"apiName"`
	ApiStatus   string    `json:"apiStatus"`
	ApiDateTime time.Time      `json:"apiDateTime"`
}