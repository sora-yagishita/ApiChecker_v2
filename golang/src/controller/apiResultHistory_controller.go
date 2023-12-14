package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/controller/dto"
	"src/model"
	"src/model/entities"
)

type ApiResultHistoryController interface {
	FetchApiResultHistory(w http.ResponseWriter, r *http.Request)
	AddApiResultHistory(w http.ResponseWriter, r *http.Request)
	UpdateApiResultHistory(w http.ResponseWriter, r *http.Request)
	DeleteApiResultHistory(w http.ResponseWriter, r *http.Request)
}

type apiResultHistoryController struct {
	apiResultHistoryModel model.ApiResultHistoryModel
}

func CreateApiResultHistoryController(apiResultHistoryModel model.ApiResultHistoryModel) ApiResultHistoryController {
	return &apiResultHistoryController{apiResultHistoryModel}
}

func (ac *apiResultHistoryController) FetchApiResultHistory(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiResultHistoryModel.FetchApiResultHistory(r)

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

func (ac *apiResultHistoryController) AddApiResultHistory(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var addApiResultHistoryRequest dto.AddApiResultHistoryRequest
	err := json.Unmarshal(body, &addApiResultHistoryRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiResultHistoryModel.AddApiResultHistory(entities.ApiResultHistory{
		ApiId:              addApiResultHistoryRequest.ApiId,
		RequestEndpoint:    addApiResultHistoryRequest.RequestEndpoint,
		RequestParamString: addApiResultHistoryRequest.RequestParamString,
		RequestDateTime:    addApiResultHistoryRequest.RequestDateTime,
		ResponseStatusCode: addApiResultHistoryRequest.ResponseStatusCode,
		ResponseData:       addApiResultHistoryRequest.ResponseData,
		ResponseDateTime:   addApiResultHistoryRequest.ResponseDateTime,
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

func (ac *apiResultHistoryController) UpdateApiResultHistory(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var updateApiResultHistoryRequest dto.UpdateApiResultHistoryRequest
	err := json.Unmarshal(body, &updateApiResultHistoryRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiResultHistoryModel.UpdateApiResultHistory(r, entities.ApiResultHistory{
		RequestEndpoint:    updateApiResultHistoryRequest.RequestEndpoint,
		RequestParamString: updateApiResultHistoryRequest.RequestParamString,
		RequestDateTime:    updateApiResultHistoryRequest.RequestDateTime,
		ResponseStatusCode: updateApiResultHistoryRequest.ResponseStatusCode,
		ResponseData:       updateApiResultHistoryRequest.ResponseData,
		ResponseDateTime:   updateApiResultHistoryRequest.ResponseDateTime,
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

func (ac *apiResultHistoryController) DeleteApiResultHistory(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiResultHistoryModel.DeleteApiResultHistory(r)

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
