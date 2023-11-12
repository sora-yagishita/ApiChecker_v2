package controller

import (
	"src/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResultController interface {
	FetchApiResult(w http.ResponseWriter, r *http.Request)
	AddApiResult(w http.ResponseWriter, r *http.Request)
	ChangeApiResult(w http.ResponseWriter, r *http.Request)
	DeleteApiResult(w http.ResponseWriter, r *http.Request)
}

type apiResultController struct {
	tm model.ApiResultModel
}

func CreateApiResultController(tm model.ApiResultModel) ApiResultController {
	return &apiResultController{tm}
}

func (tc *apiResultController) FetchApiResult(w http.ResponseWriter, r *http.Request) {
	apiResult, err := tc.tm.FetchApiResult(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(apiResult)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (tc *apiResultController) AddApiResult(w http.ResponseWriter, r *http.Request) {
	// 省略
}

func (tc *apiResultController) ChangeApiResult(w http.ResponseWriter, r *http.Request) {
	// 省略
}

func (tc *apiResultController) DeleteApiResult(w http.ResponseWriter, r *http.Request) {
	// 省略
}