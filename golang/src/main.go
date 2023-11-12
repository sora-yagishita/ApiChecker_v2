package main

import (
	"src/controller"
	"src/model"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var tm = model.CreateApiResultModel()
var tc = controller.CreateApiResultController(tm)
var ro = controller.CreateRouter(tc)


func main() {
	http.HandleFunc("/fetch-apiResult", ro.FetchApiResult)
	http.HandleFunc("/add-apiResult", ro.AddApiResult)
	http.HandleFunc("/delete-apiResult", ro.DeleteApiResult)
	http.HandleFunc("/change-apiResult", ro.ChangeApiResult)
	http.ListenAndServe(":8080", nil)
}
