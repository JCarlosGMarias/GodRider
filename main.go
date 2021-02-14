package main

import (
	"godrider/controllers"
	"godrider/infrastructures"
	"godrider/services"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

var (
	// ConfigCtrl groups all config methods
	ConfigCtrl controllers.ConfigurationControllerer
)

func init() {
	// Infrastructures
	apiUrlsIstruct := &infrastructures.APIUrlsInfrastructure{}
	var apiUrlsIstructr infrastructures.APIUrlsInfrastructurer = apiUrlsIstruct

	// Services
	configSrv := &services.ConfigurationService{}
	configSrv.APIUrlsInfrastructure(&apiUrlsIstructr)
	var configSrvr services.ConfigurationServicer = configSrv

	// Controllers
	configCtrl := controllers.ConfigurationController{}
	configCtrl.ConfigSrv(&configSrvr)
	ConfigCtrl = &configCtrl
}

func main() {
	routes := ConfigCtrl.GetRoutes()

	http.HandleFunc(routes["LoginUrl"], controllers.Login)
	// Endpoints
	http.HandleFunc(routes["GetApiUrlsUrl"], ConfigCtrl.GetApiUrls)
	// Providers
	http.HandleFunc(routes["GetProvidersUrl"], controllers.GetProviders)
	http.HandleFunc(routes["ConnectToProviderUrl"], controllers.ConnectToProvider)
	http.HandleFunc(routes["GetOrdersUrl"], controllers.GetOrders)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
