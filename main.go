package main

import (
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	diContainer := DIContainer{}
	diContainer.Register()

	routes := diContainer.ConfigCtrlr.GetRoutes()

	http.HandleFunc(routes["LoginUrl"], diContainer.UserCtrlr.Login)
	// Endpoints
	http.HandleFunc(routes["GetApiUrlsUrl"], diContainer.ConfigCtrlr.GetApiUrls)
	// Providers
	http.HandleFunc(routes["GetProvidersUrl"], diContainer.ProviderCtrlr.GetProviders)
	http.HandleFunc(routes["ConnectToProviderUrl"], diContainer.ProviderCtrlr.ConnectToProvider)
	http.HandleFunc(routes["GetOrdersUrl"], diContainer.OrderCtrlr.GetOrders)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
