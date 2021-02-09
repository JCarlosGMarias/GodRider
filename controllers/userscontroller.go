package controllers

import (
	"encoding/json"
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if helpers.IsValidMethod(w, r, []string{http.MethodPost}) {
		var userReq requests.UserRequest
		helpers.ParseBody(r.Body, &userReq)

		user, errorRs := services.UserSrv.GetUserByCredentials(&userReq)
		if errorRs == nil {
			json.NewEncoder(w).Encode(user)
		} else {
			json.NewEncoder(w).Encode(errorRs)
		}
	}
}
