package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/controller/dto"
	"src/model"
	"src/model/entities"
)

type ApiController interface {
	FetchApi(w http.ResponseWriter, r *http.Request)
	AddApi(w http.ResponseWriter, r *http.Request)
	UpdateApi(w http.ResponseWriter, r *http.Request)
	DeleteApi(w http.ResponseWriter, r *http.Request)
}

type apiController struct {
	apiModel model.ApiModel
}

func CreateApiController(apiModel model.ApiModel) ApiController {
	return &apiController{apiModel}
}

func (ac *apiController) FetchApi(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiModel.FetchApi(r)

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

func (ac *apiController) AddApi(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var addApiRequest dto.AddApiRequest
	err := json.Unmarshal(body, &addApiRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiModel.AddApi(entities.Api{
		ApiName:        addApiRequest.ApiName,
		ApiDescription: addApiRequest.ApiDescription,
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

func (ac *apiController) UpdateApi(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var updateApiRequest dto.UpdateApiRequest
	err := json.Unmarshal(body, &updateApiRequest)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	result, err := ac.apiModel.UpdateApi(r, entities.Api{
		ApiName:        updateApiRequest.ApiName,
		ApiDescription: updateApiRequest.ApiDescription,
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

func (ac *apiController) DeleteApi(w http.ResponseWriter, r *http.Request) {
	result, err := ac.apiModel.DeleteApi(r)

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
