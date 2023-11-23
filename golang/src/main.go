package main

import (
	"net/http"
	"src/controller"
	"src/model"

	_ "github.com/go-sql-driver/mysql"
)

var apiResultModel = model.CreateApiResultModel()
var apiModel = model.CreateApiModel()

var apiResultController = controller.CreateApiResultController(apiResultModel)
var apiController = controller.CreateApiController(apiModel)

var router = controller.CreateRouter(apiResultController)
var apiRouter = controller.CreateApiRouter(apiController)

func main() {
	router.HandleRequest()
	apiRouter.HandleRequest()
	http.ListenAndServe(":8080", nil)
}
