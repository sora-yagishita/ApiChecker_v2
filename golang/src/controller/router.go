package controller

import (
	"net/http"
)

type Router interface {
	HandleRequest()
}

type router struct {
	apiController              ApiController
	apiSettingController       ApiSettingController
	apiHeaderSettingController ApiHeaderSettingController
	apiParamSettingController  ApiParamSettingController
}

func CreateRouter(
	apiController ApiController,
	apiSettingController ApiSettingController,
	apiHeaderSettingController ApiHeaderSettingController,
	apiParamSettingController ApiParamSettingController) Router {
	return &router{
		apiController:              apiController,
		apiSettingController:       apiSettingController,
		apiHeaderSettingController: apiHeaderSettingController,
		apiParamSettingController:  apiParamSettingController}
}

func (ro *router) HandleRequest() {
	http.HandleFunc("/api/", ro.HandleApiRequest)
	http.HandleFunc("/api-setting/", ro.HandleApiSettingRequest)
	http.HandleFunc("/api-header-setting/", ro.HandleApiHeaderSettingRequest)
	http.HandleFunc("/api-param-setting/", ro.HandleApiParamSettingRequest)
}

func (ro *router) HandleApiRequest(w http.ResponseWriter, r *http.Request) {
	HeaderSet(w, r)
	prefix := "/api/"

	switch r.URL.Path {
	case prefix + "fetch":
		ro.apiController.FetchApi(w, r)
	case prefix + "add":
		ro.apiController.AddApi(w, r)
	case prefix + "delete":
		ro.apiController.DeleteApi(w, r)
	case prefix + "update":
		ro.apiController.UpdateApi(w, r)
	default:
		w.WriteHeader(405)
	}
}

func (ro *router) HandleApiSettingRequest(w http.ResponseWriter, r *http.Request) {
	HeaderSet(w, r)
	prefix := "/api-setting/"

	switch r.URL.Path {
	case prefix + "fetch":
		ro.apiSettingController.FetchApiSetting(w, r)
	case prefix + "add":
		ro.apiSettingController.AddApiSetting(w, r)
	case prefix + "delete":
		ro.apiSettingController.DeleteApiSetting(w, r)
	case prefix + "update":
		ro.apiSettingController.UpdateApiSetting(w, r)
	default:
		w.WriteHeader(405)
	}
}

func (ro *router) HandleApiHeaderSettingRequest(w http.ResponseWriter, r *http.Request) {
	HeaderSet(w, r)
	prefix := "/api-header-setting/"

	switch r.URL.Path {
	case prefix + "fetch":
		ro.apiHeaderSettingController.FetchApiHeaderSetting(w, r)
	case prefix + "add":
		ro.apiHeaderSettingController.AddApiHeaderSetting(w, r)
	case prefix + "delete":
		ro.apiHeaderSettingController.DeleteApiHeaderSetting(w, r)
	case prefix + "update":
		ro.apiHeaderSettingController.UpdateApiHeaderSetting(w, r)
	default:
		w.WriteHeader(405)
	}
}

func (ro *router) HandleApiParamSettingRequest(w http.ResponseWriter, r *http.Request) {
	HeaderSet(w, r)
	prefix := "/api-param-setting/"

	switch r.URL.Path {
	case prefix + "fetch":
		ro.apiParamSettingController.FetchApiParamSetting(w, r)
	case prefix + "add":
		ro.apiParamSettingController.AddApiParamSetting(w, r)
	case prefix + "delete":
		ro.apiParamSettingController.DeleteApiParamSetting(w, r)
	case prefix + "update":
		ro.apiParamSettingController.UpdateApiParamSetting(w, r)
	default:
		w.WriteHeader(405)
	}
}

func HeaderSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}
}
