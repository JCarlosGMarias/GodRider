package controllers

import (
	"encoding/json"
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

func GetProviders(w http.ResponseWriter, r *http.Request) {
	if helpers.IsValidMethod(w, r, []string{"POST"}) {
		var providerRq requests.ProviderRequest
		helpers.ParseBody(r.Body, &providerRq)

		if helpers.IsValidToken(w, r, providerRq.Token) {
			json.NewEncoder(w).Encode(services.ProviderSrv.GetAllProviders())
		}
	}
}
