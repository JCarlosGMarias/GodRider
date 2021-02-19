package main

import (
	"godrider/controllers"
	"godrider/infrastructures"
	"godrider/services"
	"godrider/webclients"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

var (
	// ConfigCtrlr groups all config methods
	ConfigCtrlr controllers.ConfigurationControllerer
	// UserCtrlr groups all user main methods, as the Login method for authentication into the API
	UserCtrlr controllers.UserControllerer
	// ProviderCtrlr groups all provider main methods
	ProviderCtrlr controllers.ProviderControllerer
	// OrderCtrlr groups all order-involved methods
	OrderCtrlr controllers.OrderControllerer
)

func init() {
	// Infrastructures
	userProviderIstruct := &infrastructures.UserProviderInfrastructure{}
	userProviderIstruct.TableName("userprovider")

	var apiUrlsIstructr infrastructures.APIUrlsInfrastructurer = &infrastructures.APIUrlsInfrastructure{}
	var userIstructr infrastructures.UserInfrastructurer = &infrastructures.UserInfrastructure{}
	var userProviderIstructr infrastructures.UserProviderInfrastructurer = userProviderIstruct
	var providerIstructr infrastructures.ProviderInfrastructurer = &infrastructures.ProviderInfrastructure{}
	var webClientFactorier webclients.WebClientFactorier = &webclients.ClientFactory

	// Services
	configSrv := &services.ConfigurationService{}
	userSrv := &services.UserService{}
	providerSrv := &services.ProviderService{}
	userProviderSrv := &services.UserProviderService{}
	orderSrv := &services.OrderService{}
	validationSrv := &services.ValidationService{}

	configSrv.APIUrlsInfrastructure(&apiUrlsIstructr)
	userSrv.UserInfrastructure(&userIstructr)
	providerSrv.ProviderInfrastructure(&providerIstructr)
	userProviderSrv.UserProviderInfrastructure(&userProviderIstructr)
	orderSrv.ProviderInfrastructure(&providerIstructr)
	orderSrv.Factory(&webClientFactorier)

	var configSrvr services.ConfigurationServicer = configSrv
	var userSrvr services.UserServicer = userSrv
	var providerSrvr services.ProviderServicer = providerSrv
	var userProviderSrvr services.UserProviderServicer = userProviderSrv
	var orderSrvr services.OrderServicer = orderSrv

	validationSrv.UserSrv(&userSrvr)
	var validationSrvr services.ValidationServicer = validationSrv

	// Controllers
	configCtrl := &controllers.ConfigurationController{}
	userCtrl := &controllers.UserController{}
	providerCtrl := &controllers.ProviderController{}
	orderCtrl := &controllers.OrderController{}

	configCtrl.ConfigSrv(&configSrvr)
	configCtrl.ValidationSrv(&validationSrvr)
	userCtrl.UserSrv(&userSrvr)
	userCtrl.ValidationSrv(&validationSrvr)
	providerCtrl.ProviderSrv(&providerSrvr)
	providerCtrl.UserProviderSrv(&userProviderSrvr)
	providerCtrl.ValidationSrv(&validationSrvr)
	orderCtrl.OrderSrv(&orderSrvr)
	orderCtrl.ValidationSrv(&validationSrvr)

	ConfigCtrlr = configCtrl
	UserCtrlr = userCtrl
	ProviderCtrlr = providerCtrl
	OrderCtrlr = orderCtrl
}

func main() {
	routes := ConfigCtrlr.GetRoutes()

	http.HandleFunc(routes["LoginUrl"], UserCtrlr.Login)
	// Endpoints
	http.HandleFunc(routes["GetApiUrlsUrl"], ConfigCtrlr.GetApiUrls)
	// Providers
	http.HandleFunc(routes["GetProvidersUrl"], ProviderCtrlr.GetProviders)
	http.HandleFunc(routes["ConnectToProviderUrl"], ProviderCtrlr.ConnectToProvider)
	http.HandleFunc(routes["GetOrdersUrl"], OrderCtrlr.GetOrders)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
