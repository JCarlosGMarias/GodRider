package controllers

import (
	"encoding/json"
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

func GetRoutes() map[string]string {
	return services.ConfigSrv.GetApiUrls()
}

func GetApiUrls(w http.ResponseWriter, r *http.Request) {
	if helpers.IsValidMethod(w, r, []string{"POST"}) {
		var apiUrlsRq requests.ApiUrlsRequest
		helpers.ParseBody(r.Body, &apiUrlsRq)

		if helpers.IsValidToken(w, r, apiUrlsRq.Token) {
			json.NewEncoder(w).Encode(services.ConfigSrv.GetApiUrls())
		}
	}
}
