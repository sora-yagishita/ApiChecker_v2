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
	router.HandleRequest()
	http.ListenAndServe(":8080", nil)
}
