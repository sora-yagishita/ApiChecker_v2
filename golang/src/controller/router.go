package controller

import (
	"net/http"
)

type Router interface {
	HandleRequest()
}

type router struct {
	apiResultController ApiResultController
}

func CreateRouter(apiResultController ApiResultController) Router {
	return &router{apiResultController}
}

func (ro *router) HandleRequest() {
	http.HandleFunc("/api-result/", ro.HandleApiResultRequest)
}

func (ro *router) HandleApiResultRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	prefix := "/api-result/"

	switch r.URL.Path {
	case prefix + "fetch-apiResult":
		ro.apiResultController.FetchApiResult(w, r)
	case prefix + "add-apiResult":
		ro.apiResultController.AddApiResult(w, r)
	case prefix + "delete-apiResult":
		ro.apiResultController.DeleteApiResult(w, r)
	case prefix + "change-apiResult":
		ro.apiResultController.ChangeApiResult(w, r)
	default:
		w.WriteHeader(405)
	}
}
