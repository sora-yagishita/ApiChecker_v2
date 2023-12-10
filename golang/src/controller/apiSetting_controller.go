package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/controller/dto"
	"src/model"
	"src/model/entities"
)

type ApiSettingController interface {
	FetchApiSetting(w http.ResponseWriter, r *http.Request)
	AddApiSetting(w http.ResponseWriter, r *http.Request)
	UpdateApiSetting(w http.ResponseWriter, r *http.Request)
	DeleteApiSetting(w http.ResponseWriter, r *http.Request)
}

type apiSettingController struct {
	apiSettingModel model.ApiSettingModel
}

func CreateApiSettingController(apiSettingModel model.ApiSettingModel) ApiSettingController {
	return &apiSettingController{apiSettingModel}
}

func (ac *apiSettingController) FetchApiSetting(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiSettingModel.FetchApiSetting(r)

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

func (ac *apiSettingController) AddApiSetting(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var addApiSettingRequest dto.AddApiSettingRequest
	err := json.Unmarshal(body, &addApiSettingRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiSettingModel.AddApiSetting(entities.ApiSetting{
		ApiId:                addApiSettingRequest.ApiId,
		RequestMethod:        addApiSettingRequest.RequestMethod,
		Endpoint:             addApiSettingRequest.Endpoint,
		ExecutionIntervalSec: addApiSettingRequest.ExecutionIntervalSec,
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

func (ac *apiSettingController) UpdateApiSetting(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var updateApiSettingRequest dto.UpdateApiSettingRequest
	err := json.Unmarshal(body, &updateApiSettingRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiSettingModel.UpdateApiSetting(r, entities.ApiSetting{
		RequestMethod:        updateApiSettingRequest.RequestMethod,
		Endpoint:             updateApiSettingRequest.Endpoint,
		ExecutionIntervalSec: updateApiSettingRequest.ExecutionIntervalSec,
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

func (ac *apiSettingController) DeleteApiSetting(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiSettingModel.DeleteApiSetting(r)

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
