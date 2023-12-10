package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/controller/dto"
	"src/model"
	"src/model/entities"
)

type ApiParamSettingController interface {
	FetchApiParamSetting(w http.ResponseWriter, r *http.Request)
	AddApiParamSetting(w http.ResponseWriter, r *http.Request)
	UpdateApiParamSetting(w http.ResponseWriter, r *http.Request)
	DeleteApiParamSetting(w http.ResponseWriter, r *http.Request)
}

type apiParamSettingController struct {
	apiParamSettingModel model.ApiParamSettingModel
}

func CreateApiParamSettingController(apiParamSettingModel model.ApiParamSettingModel) ApiParamSettingController {
	return &apiParamSettingController{apiParamSettingModel}
}

func (ac *apiParamSettingController) FetchApiParamSetting(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiParamSettingModel.FetchApiParamSetting(r)

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

func (ac *apiParamSettingController) AddApiParamSetting(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var addApiParamSettingRequest dto.AddApiParamSettingRequest
	err := json.Unmarshal(body, &addApiParamSettingRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiParamSettingModel.AddApiParamSetting(entities.ApiParamSetting{
		ApiSettingId:      addApiParamSettingRequest.ApiSettingId,
		ApiParamSettingNo: addApiParamSettingRequest.ApiParamSettingNo,
		ApiParamKey:       addApiParamSettingRequest.ApiParamKey,
		ApiParamValue:     addApiParamSettingRequest.ApiParamValue,
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

func (ac *apiParamSettingController) UpdateApiParamSetting(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var updateApiParamSettingRequest dto.UpdateApiParamSettingRequest
	err := json.Unmarshal(body, &updateApiParamSettingRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiParamSettingModel.UpdateApiParamSetting(r, entities.ApiParamSetting{
		ApiParamKey:   updateApiParamSettingRequest.ApiParamKey,
		ApiParamValue: updateApiParamSettingRequest.ApiParamValue,
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

func (ac *apiParamSettingController) DeleteApiParamSetting(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiParamSettingModel.DeleteApiParamSetting(r)

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
