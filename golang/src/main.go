package main

import (
	"src/controller"
	"src/model"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var apiResultModel = model.CreateApiResultModel()
var apiResultController = controller.CreateApiResultController(apiResultModel)
var router = controller.CreateRouter(apiResultController)


func main() {
	// http.HandleFunc("/fetch-apiResult", router.FetchApiResult)
	// http.HandleFunc("/add-apiResult", router.AddApiResult)
	// http.HandleFunc("/delete-apiResult", router.DeleteApiResult)
	// http.HandleFunc("/change-apiResult", router.ChangeApiResult)
	router.HandleRequest()
	http.ListenAndServe(":8080", nil)
}
