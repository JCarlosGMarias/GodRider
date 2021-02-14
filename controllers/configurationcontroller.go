package controllers

import (
	"encoding/json"
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

type ConfigurationControllerer interface {
	GetRoutes() map[string]string
	GetApiUrls(w http.ResponseWriter, r *http.Request)
}

type ConfigurationController struct {
	configSrv services.ConfigurationServicer
}

func (ctrl *ConfigurationController) GetRoutes() map[string]string {
	return ctrl.configSrv.GetAPIUrls()
}

func (ctrl *ConfigurationController) GetApiUrls(w http.ResponseWriter, r *http.Request) {
	if helpers.IsValidMethod(w, r, []string{http.MethodPost}) {
		var apiUrlsRq requests.ApiUrlsRequest
		helpers.ParseBody(r.Body, &apiUrlsRq)

		if helpers.IsValidToken(w, r, apiUrlsRq.Token) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ctrl.configSrv.GetAPIUrls())
		}
	}
}

// ConfigSrv setter
func (ctrl *ConfigurationController) ConfigSrv(service *services.ConfigurationServicer) {
	if ctrl.configSrv == nil {
		ctrl.configSrv = *service
	}
}
