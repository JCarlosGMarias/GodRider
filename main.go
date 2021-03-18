package main

import (
	"godrider/helpers"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func handleOptions(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			helpers.SetOptionsResponseEncoder(&w)
			return
		}

		f(w, r)
	}
}

func main() {
	diContainer := DIContainer{}
	diContainer.Register()

	routes := diContainer.ConfigCtrlr.GetRoutes()

	http.HandleFunc(routes["LoginUrl"], handleOptions(diContainer.UserCtrlr.Login))
	// Endpoints
	http.HandleFunc(routes["GetApiUrlsUrl"], handleOptions(diContainer.ConfigCtrlr.GetApiUrls))
	// Providers
	http.HandleFunc(routes["GetProvidersUrl"], handleOptions(diContainer.ProviderCtrlr.GetProviders))
	http.HandleFunc(routes["ConnectToProviderUrl"], handleOptions(diContainer.ProviderCtrlr.ConnectToProvider))
	http.HandleFunc(routes["GetOrdersUrl"], handleOptions(diContainer.OrderCtrlr.GetOrders))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
