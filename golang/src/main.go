package main

import (
	"net/http"
	"src/controller"
	"src/model"

	_ "github.com/go-sql-driver/mysql"
)

var apiModel = model.CreateApiModel()
var apiSettingModel = model.CreateApiSettingModel()
var apiHeaderSettingModel = model.CreateApiHeaderSettingModel()
var apiParamSettingModel = model.CreateApiParamSettingModel()

var apiController = controller.CreateApiController(apiModel)
var apiSettingController = controller.CreateApiSettingController(apiSettingModel)
var apiHeaderSettingController = controller.CreateApiHeaderSettingController(apiHeaderSettingModel)
var apiParamSettingController = controller.CreateApiParamSettingController(apiParamSettingModel)

var router = controller.CreateRouter(
	apiController,
	apiSettingController,
	apiHeaderSettingController,
	apiParamSettingController)

func main() {
	router.HandleRequest()
	http.ListenAndServe(":8080", nil)
}
