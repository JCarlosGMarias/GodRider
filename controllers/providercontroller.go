package controllers

import (
	"encoding/json"
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

type ProviderControllerer interface {
	ConnectToProvider(w http.ResponseWriter, r *http.Request)
}

type ProviderController struct {
	userProviderSrv services.UserProviderServicer
}

func GetProviders(w http.ResponseWriter, r *http.Request) {
	if helpers.IsValidMethod(w, r, []string{http.MethodPost}) {
		var providerRq requests.ProviderRequest
		helpers.ParseBody(r.Body, &providerRq)

		if helpers.IsValidToken(w, r, providerRq.Token) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(services.ProviderSrv.GetAllProviders())
		}
	}
}

func (ctrl *ProviderController) ConnectToProvider(w http.ResponseWriter, r *http.Request) {
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

			var err error
			if r.Method == http.MethodPost {
				err = ctrl.userProviderSrv.AddConnection(&request)
			} else if r.Method == http.MethodPut {
				err = ctrl.userProviderSrv.UpdateConnection(&request)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
		}
	}
}

// UserProviderSrv setter
func (ctrl *ProviderController) UserProviderSrv(srv *services.UserProviderServicer) {
	if ctrl.userProviderSrv == nil {
		ctrl.userProviderSrv = *srv
	}
}
