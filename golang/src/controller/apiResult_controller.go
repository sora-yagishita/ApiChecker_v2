package controller

import (
	"src/controller/dto"
	"src/model/entities"
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
	apiResultModel model.ApiResultModel
}

func CreateApiResultController(apiResultModel model.ApiResultModel) ApiResultController {
	return &apiResultController{apiResultModel}
}

func (ac *apiResultController) FetchApiResult(w http.ResponseWriter, r *http.Request) {
	apiResult, err := ac.apiResultModel.FetchApiResult(r)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(apiResult)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (ac *apiResultController) AddApiResult(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var addApiResultRequest dto.AddApiResultRequest
	err := json.Unmarshal(body, &addApiResultRequest)

	fmt.Println("表示0")
	fmt.Println(addApiResultRequest.ApiName, addApiResultRequest.ApiStatus, addApiResultRequest.ApiDateTime)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	fmt.Println("表示")
	fmt.Println(addApiResultRequest.ApiName, addApiResultRequest.ApiStatus, addApiResultRequest.ApiDateTime)

	result, err := ac.apiResultModel.AddApiResult(entities.ApiResult{
		ApiName: addApiResultRequest.ApiName,
		ApiStatus: addApiResultRequest.ApiStatus,
		ApiDateTime: addApiResultRequest.ApiDateTime,
	})

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (tc *apiResultController) ChangeApiResult(w http.ResponseWriter, r *http.Request) {
	// 省略
}

func (tc *apiResultController) DeleteApiResult(w http.ResponseWriter, r *http.Request) {
	// 省略
}