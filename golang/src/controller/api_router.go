package controller

import (
	"net/http"
)

type ApiRouter interface {
	HandleRequest()
}

type apiRouter struct {
	apiController ApiController
}

func CreateApiRouter(apiController ApiController) ApiRouter {
	return &apiRouter{apiController}
}

func (ro *apiRouter) HandleRequest() {
	http.HandleFunc("/api/", ro.HandleApiRequest)
}

func (ro *apiRouter) HandleApiRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	prefix := "/api/"

	switch r.URL.Path {
	case prefix + "fetch-api":
		ro.apiController.FetchApi(w, r)
	case prefix + "add-api":
		ro.apiController.AddApi(w, r)
	case prefix + "delete-api":
		ro.apiController.DeleteApi(w, r)
	case prefix + "update-api":
		ro.apiController.UpdateApi(w, r)
	default:
		w.WriteHeader(405)
	}
}
