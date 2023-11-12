package controller

import (
	"net/http"
)

type Router interface {
	FetchApiResult(w http.ResponseWriter, r *http.Request)
	AddApiResult(w http.ResponseWriter, r *http.Request)
	DeleteApiResult(w http.ResponseWriter, r *http.Request)
	ChangeApiResult(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc ApiResultController
}

func CreateRouter(tc ApiResultController) Router {
	return &router{tc}
}

func (ro *router) FetchApiResult(w http.ResponseWriter, r *http.Request) {
	// プリフライトリクエスト用に設定している。
	w.Header().Set("Access-Control-Allow-Headers", "*")

    // CORSエラー対策。APIを叩くフロント側のURLを渡す。
	w.Header().Set("Access-Control-Allow-Origin", "*")

    // 返却する値のContent-Typeを設定。
	w.Header().Set("Content-Type", "application/json")

	// preflightでAPIが二度実行されてしまうことを防ぐ。
	if r.Method == "OPTIONS" {
		return
	}

	// controllerを呼び出す。
	ro.tc.FetchApiResult(w, r)
}

func (ro *router) AddApiResult(w http.ResponseWriter, r *http.Request) {
	// 省略
}

func (ro *router) DeleteApiResult(w http.ResponseWriter, r *http.Request) {
	// 省略
}

func (ro *router) ChangeApiResult(w http.ResponseWriter, r *http.Request) {
	// 省略
}
