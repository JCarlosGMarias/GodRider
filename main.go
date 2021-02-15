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
	// ProviderCtrl groups all provider main methods
	ProviderCtrl controllers.ProviderControllerer
)

func init() {
	// Infrastructures
	apiUrlsIstruct := &infrastructures.APIUrlsInfrastructure{}
	var apiUrlsIstructr infrastructures.APIUrlsInfrastructurer = apiUrlsIstruct
	userProviderIstruct := &infrastructures.UserProviderInfrastructure{}
	userProviderIstruct.TableName("userprovider")
	var userProviderIstructr infrastructures.UserProviderInfrastructurer = userProviderIstruct

	// Services
	configSrv := &services.ConfigurationService{}
	configSrv.APIUrlsInfrastructure(&apiUrlsIstructr)
	var configSrvr services.ConfigurationServicer = configSrv
	userProviderSrv := &services.UserProviderService{}
	userProviderSrv.UserProviderInfrastructure(&userProviderIstructr)
	var userProviderSrvr services.UserProviderServicer = userProviderSrv

	// Controllers
	configCtrl := &controllers.ConfigurationController{}
	configCtrl.ConfigSrv(&configSrvr)
	ConfigCtrl = configCtrl
	providerCtrl := &controllers.ProviderController{}
	providerCtrl.UserProviderSrv(&userProviderSrvr)
	ProviderCtrl = providerCtrl
}

func main() {
	routes := ConfigCtrl.GetRoutes()

	http.HandleFunc(routes["LoginUrl"], controllers.Login)
	// Endpoints
	http.HandleFunc(routes["GetApiUrlsUrl"], ConfigCtrl.GetApiUrls)
	// Providers
	http.HandleFunc(routes["GetProvidersUrl"], controllers.GetProviders)
	http.HandleFunc(routes["ConnectToProviderUrl"], ProviderCtrl.ConnectToProvider)
	http.HandleFunc(routes["GetOrdersUrl"], controllers.GetOrders)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
