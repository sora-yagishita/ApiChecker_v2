package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/controller/dto"
	"src/model"
	"src/model/entities"
)

type ApiHeaderSettingController interface {
	FetchApiHeaderSetting(w http.ResponseWriter, r *http.Request)
	AddApiHeaderSetting(w http.ResponseWriter, r *http.Request)
	UpdateApiHeaderSetting(w http.ResponseWriter, r *http.Request)
	DeleteApiHeaderSetting(w http.ResponseWriter, r *http.Request)
}

type apiHeaderSettingController struct {
	apiHeaderSettingModel model.ApiHeaderSettingModel
}

func CreateApiHeaderSettingController(apiHeaderSettingModel model.ApiHeaderSettingModel) ApiHeaderSettingController {
	return &apiHeaderSettingController{apiHeaderSettingModel}
}

func (ac *apiHeaderSettingController) FetchApiHeaderSetting(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiHeaderSettingModel.FetchApiHeaderSetting(r)

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

func (ac *apiHeaderSettingController) AddApiHeaderSetting(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var addApiHeaderSettingRequest dto.AddApiHeaderSettingRequest
	err := json.Unmarshal(body, &addApiHeaderSettingRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiHeaderSettingModel.AddApiHeaderSetting(entities.ApiHeaderSetting{
		ApiSettingId:       addApiHeaderSettingRequest.ApiSettingId,
		ApiHeaderSettingNo: addApiHeaderSettingRequest.ApiHeaderSettingNo,
		ApiHeaderKey:       addApiHeaderSettingRequest.ApiHeaderKey,
		ApiHeaderValue:     addApiHeaderSettingRequest.ApiHeaderValue,
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

func (ac *apiHeaderSettingController) UpdateApiHeaderSetting(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var updateApiHeaderSettingRequest dto.UpdateApiHeaderSettingRequest
	err := json.Unmarshal(body, &updateApiHeaderSettingRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiHeaderSettingModel.UpdateApiHeaderSetting(r, entities.ApiHeaderSetting{
		ApiHeaderKey:   updateApiHeaderSettingRequest.ApiHeaderKey,
		ApiHeaderValue: updateApiHeaderSettingRequest.ApiHeaderValue,
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

func (ac *apiHeaderSettingController) DeleteApiHeaderSetting(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiHeaderSettingModel.DeleteApiHeaderSetting(r)

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
