package main

import (
	"log"
	"net/http"

	"godrider/controllers"

	_ "modernc.org/sqlite"
)

func main() {
	routes := controllers.GetRoutes()

	http.HandleFunc(routes["LoginUrl"], controllers.Login)
	// Endpoints
	http.HandleFunc(routes["GetApiUrlsUrl"], controllers.GetApiUrls)
	// Providers
	http.HandleFunc(routes["GetProvidersUrl"], controllers.GetProviders)
	http.HandleFunc(routes["ConnectToProviderUrl"], controllers.ConnectToProvider)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
