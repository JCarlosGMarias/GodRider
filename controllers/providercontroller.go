package controllers

import (
	"encoding/json"
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

func GetProviders(w http.ResponseWriter, r *http.Request) {
	if helpers.IsValidMethod(w, r, []string{http.MethodPost}) {
		var providerRq requests.ProviderRequest
		helpers.ParseBody(r.Body, &providerRq)

		if helpers.IsValidToken(w, r, providerRq.Token) {
			json.NewEncoder(w).Encode(services.ProviderSrv.GetAllProviders())
		}
	}
}

func ConnectToProvider(w http.ResponseWriter, r *http.Request) {
	if helpers.IsValidMethod(w, r, []string{http.MethodPost, http.MethodPut}) {
		var providerRq requests.ProviderRequest
		helpers.ParseBody(r.Body, &providerRq)

		if helpers.IsValidToken(w, r, providerRq.Token) {
			user, _ := services.UserSrv.GetUserByToken(&requests.UserRequest{Token: providerRq.Token})

			request := requests.UserProviderRequest{
				UserId:     user.ID,
				ProviderId: providerRq.ProviderID,
				IsActive:   providerRq.IsActive,
			}

			if r.Method == http.MethodPost {
				err := services.UserProviderSrv.AddConnection(&request)
				json.NewEncoder(w).Encode(err)
			} else if r.Method == http.MethodPut {
				err := services.UserProviderSrv.UpdateConnection(&request)
				json.NewEncoder(w).Encode(err)
			}
		}
	}
}
